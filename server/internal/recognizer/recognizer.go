package recognizer

import (
	"animal-sound-recognizer/internal/file_storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var recognizerUrl = os.Getenv("RECOGNIZER_URL")

func ProcessAudio(fileId string) (RecognizeResult, error) {
	fileBytes, fileName, err := file_storage.GetFile(fileId)
	if err != nil {
		return RecognizeResult{}, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return RecognizeResult{}, fmt.Errorf("Ошибка создания части form-data: %v\n", err)
	}
	_, err = io.Copy(part, bytes.NewReader(fileBytes))
	if err != nil {
		return RecognizeResult{}, fmt.Errorf("Ошибка копирования данных файла: %v\n", err)
	}

	err = writer.WriteField("confidence_threshold", "0.5")
	if err != nil {
		return RecognizeResult{}, fmt.Errorf("Error writing field: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return RecognizeResult{}, fmt.Errorf("Error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", recognizerUrl, body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return RecognizeResult{}, fmt.Errorf("Error sending request:", err)
	}
	defer resp.Body.Close()

	var response RecognizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return RecognizeResult{}, fmt.Errorf(err.Error())
	}

	return recognizeResponseToRecognizeResult(response), nil
}
