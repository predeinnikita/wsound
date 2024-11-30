package recognizer

import (
	"animal-sound-recognizer/internal/file_storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

const recognizerUrl = "http://localhost:8000/process-audio/"

func ProcessAudio(fileId string) RecognizeResult {
	fileBytes, fileName, err := file_storage.GetFile(fileId)
	if err != nil {
		return RecognizeResult{}
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName) // Используем имя файла из file_storage
	if err != nil {
		fmt.Printf("Ошибка создания части form-data: %v\n", err)
		return RecognizeResult{}
	}
	_, err = io.Copy(part, bytes.NewReader(fileBytes))
	if err != nil {
		fmt.Printf("Ошибка копирования данных файла: %v\n", err)
		return RecognizeResult{}
	}

	err = writer.WriteField("confidence_threshold", "0.5")
	if err != nil {
		log.Fatalf("Error writing field: %v", err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatalf("Error closing writer: %v", err)
	}

	req, err := http.NewRequest("POST", recognizerUrl, body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return RecognizeResult{}
	}
	defer resp.Body.Close()

	var response RecognizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}

	return recognizeResponseToRecognizeResult(response)
}
