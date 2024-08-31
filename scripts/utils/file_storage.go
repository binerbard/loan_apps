package utils

import (
	"context"
	"loan_apps/config/setting"
	"loan_apps/internal/infrastructure/storage"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func UploadImage(file *multipart.FileHeader)(string, error) {
	bucketName := setting.StorageBucket
	objectName := file.Filename

	f,err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Upload the file to specific bucket
	minioClient := storage.NewStorage()
	// Upload file ke MinIO
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, f, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}

	fileURL := "http://localhost:9000/" + bucketName + "/" + objectName

	return fileURL, nil
}