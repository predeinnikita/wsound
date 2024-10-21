package file_storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

var endpoint = "localhost:9000"
var accessKeyID = "bIfSnMq0spOLvyXX1Zhu"
var secretAccessKey = "bqNeN1gi3Bz5b7M4qi5CbsM7JDqacsElxCguudqW"
var useSSL = false

var bucketName = "file"

var minioClient, _ = minio.New(endpoint, &minio.Options{
	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	Secure: useSSL,
})

func SaveFile(filename string, file []byte) (string, error) {
	contentType := "application/octet-stream"
	ctx := context.Background()

	fmt.Println(1)
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			log.Fatalln(err)
		}
	}

	fileReader := bytes.NewReader(file)
	fileSize := int64(len(file))

	id := uuid.New().String()

	info, err := minioClient.PutObject(ctx, bucketName, id, fileReader, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
		UserMetadata: map[string]string{
			"Filename": filename,
		},
	})
	if err != nil {
		return "", err
	}

	log.Printf("Файл успешно загружен: %s (размер: %d байт)", info.Key, info.Size)
	return id, nil
}

func DeleteFile(id string) error {
	ctx := context.Background()

	err := minioClient.RemoveObject(ctx, bucketName, id, minio.RemoveObjectOptions{})
	if err != nil {
		log.Println("Ошибка при удалении файла:", err)
		return err
	}

	log.Printf("Файл %s успешно удалён", id)
	return nil
}

func GetFile(id string) ([]byte, string, error) {
	ctx := context.Background()

	object, err := minioClient.GetObject(ctx, bucketName, id, minio.GetObjectOptions{})
	defer object.Close()

	if err != nil {
		log.Println("Ошибка при получении файла:", err)
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, object)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		return nil, "", err
	}

	stat, err := object.Stat()
	if err != nil {
		log.Println("Ошибка при получении метаданных файла:", err)
		return nil, "", err
	}

	originalFilename := stat.UserMetadata["Filename"]
	if originalFilename == "" {
		originalFilename = id
	}

	log.Printf("Файл %s успешно получен", originalFilename)
	return buf.Bytes(), originalFilename, nil
}
