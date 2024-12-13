package file_storage

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

var endpoint = os.Getenv("MINIO_ENDPOINT")
var accessKeyID = os.Getenv("MINIO_ACCESS_KEY_ID")
var secretAccessKey = os.Getenv("MINIO_SECRET_ACCESS_KEY")
var useSSL = os.Getenv("MINIO_USE_SSL") == "true"
var bucketName = os.Getenv("MINIO_BUCKET_NAME")

var minioClient, _ = minio.New(endpoint, &minio.Options{
	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	Secure: useSSL,
})

func SaveFile(filename string, file []byte) (string, error) {
	contentType := "application/octet-stream"
	ctx := context.Background()
	createBucketIfNotExists(ctx)

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
	createBucketIfNotExists(ctx)

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
	createBucketIfNotExists(ctx)

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

func createBucketIfNotExists(ctx context.Context) {
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
}
