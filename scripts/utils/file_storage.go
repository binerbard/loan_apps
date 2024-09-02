package utils

import (
	"bytes"
	"context"
	"loan_apps/config/setting"
	"loan_apps/internal/infrastructure/storage"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

var minioClient = storage.NewStorage()

func UploadImage(file *multipart.FileHeader)(string, error) {
	bucketName := setting.StorageBucket
	objectName := uuid.New().String() + "." + strings.Split(file.Filename, ".")[1]

	f,err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Upload the file to specific bucket
	// Upload file ke MinIO
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, f, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}

	return objectName, nil
}


func ShowFile(objectReader string) (*bytes.Reader, error) {
	bucketName := setting.StorageBucket
	objectName := objectReader
	object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(object)
	if err != nil {
		return nil, err
	}

	objectResponse := bytes.NewReader(buf.Bytes())
	return objectResponse, nil
}