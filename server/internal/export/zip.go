package export

import (
	"animal-sound-recognizer/internal/audio"
	"animal-sound-recognizer/internal/file_storage"
	"animal-sound-recognizer/internal/projects"
	"archive/zip"
	"bytes"
	"fmt"
	goaudio "github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"os"
	"strconv"
	"strings"
)

func GetZipArchiveIntervalsByAudio(audioId uint64) (*os.File, error) {
	audioEntity, err := audio.GetAudio(audioId)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	audioFiles := make([][]byte, 0)
	audioFileNames := make([]string, 0)

	if len(audioEntity.Intervals) > 0 {
		audioFile, _, _ := file_storage.GetFile(audioEntity.StorageID)
		intervals, _ := ExtractIntervalsFromWav(audioFile, audioEntity.Intervals)
		for i, interval := range intervals {
			audioFiles = append(audioFiles, interval)
			audioFileNames = append(audioFileNames, fmt.Sprintf("%d_%s", i+1, audioEntity.Name))
		}
	}

	archive, err := createZipArchive(
		fmt.Sprintf("%s.zip", audioEntity.Name),
		audioFiles,
		audioFileNames,
	)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return archive, nil
}

func GetZipArchiveIntervalsByProject(projectId uint64) (*os.File, error) {
	audios, err := audio.GetAllAudios(projectId)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	project, err := projects.GetProject(projectId)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	audioFiles := make([][]byte, 0)
	audioFileNames := make([]string, 0)
	for _, audioEntity := range audios {
		if len(audioEntity.Intervals) == 0 {
			continue
		}
		audioFile, fileName, _ := file_storage.GetFile(audioEntity.StorageID)
		intervals, _ := ExtractIntervalsFromWav(audioFile, audioEntity.Intervals)
		for i, interval := range intervals {
			audioFiles = append(audioFiles, interval)
			audioFileNames = append(audioFileNames, fmt.Sprintf("%s_%d_%s", project.Name, i+1, fileName))
		}
	}

	archive, err := createZipArchive(
		fmt.Sprintf("%s.zip", project.Name),
		audioFiles,
		audioFileNames,
	)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return archive, nil
}

func createZipArchive(zipFileName string, files [][]byte, fileNames []string) (*os.File, error) {
	if len(files) != len(fileNames) {
		return nil, fmt.Errorf("количество файлов и имён файлов не совпадает")
	}

	tempFile, err := os.Create(zipFileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания zip файла: %w", err)
	}

	zipWriter := zip.NewWriter(tempFile)

	for i, fileContent := range files {
		fileWriter, err := zipWriter.Create(fileNames[i])
		if err != nil {
			zipWriter.Close()
			tempFile.Close()
			os.Remove(tempFile.Name())
			return nil, fmt.Errorf("ошибка создания файла %s в архиве: %w", fileNames[i], err)
		}

		_, err = fileWriter.Write(fileContent)
		if err != nil {
			zipWriter.Close()
			tempFile.Close()
			os.Remove(tempFile.Name())
			return nil, fmt.Errorf("ошибка записи данных в файл %s: %w", fileNames[i], err)
		}
	}

	if err := zipWriter.Close(); err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("ошибка закрытия zip файла: %w", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("ошибка перемотки файла: %w", err)
	}

	return tempFile, nil
}

func ExtractIntervalsFromWav(wavData []byte, intervals []audio.Interval) ([][]byte, error) {
	decoder := wav.NewDecoder(bytes.NewReader(wavData))
	decoded, _ := decoder.FullPCMBuffer()
	sampleRate := decoded.Format.SampleRate
	numChannels := decoded.Format.NumChannels
	bitDepth := decoder.BitDepth

	var results [][]byte

	for _, interval := range intervals {
		startSec, err := parseTimeToSeconds(interval.Start)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга времени Start: %w", err)
		}
		endSec, err := parseTimeToSeconds(interval.End)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга времени End: %w", err)
		}

		startSample := int(startSec * float64(sampleRate))
		endSample := int(endSec * float64(sampleRate))

		intervalBuffer := &goaudio.IntBuffer{
			Format:         decoder.Format(),
			Data:           decoded.Data[startSample*numChannels : min(endSample*numChannels, len(decoded.Data))],
			SourceBitDepth: decoded.SourceBitDepth,
		}

		out, err := os.Create("tmp_out.wav")
		if err != nil {
			return nil, err
		}
		writer := wav.NewEncoder(out, sampleRate, int(bitDepth), numChannels, 1)
		if err := writer.Write(intervalBuffer); err != nil {
			return nil, fmt.Errorf("ошибка записи WAV интервала: %w", err)
		}
		if err := writer.Close(); err != nil {
			return nil, fmt.Errorf("ошибка закрытия WAV writer: %w", err)
		}

		res, err := os.ReadFile("tmp_out.wav")
		results = append(results, res)
		os.RemoveAll("tmp_out.wav")
	}

	return results, nil
}

func parseTimeToSeconds(timeStr string) (float64, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат времени: %s", timeStr)
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга часов: %w", err)
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга минут: %w", err)
	}

	seconds, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга секунд: %w", err)
	}

	totalSeconds := float64(hours*3600+minutes*60) + seconds
	return totalSeconds, nil
}
