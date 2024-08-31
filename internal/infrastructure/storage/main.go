package storage

import (
	"context"
	"loan_apps/config/setting"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// NewStorage menginisialisasi MinIO client dan memastikan bucket ada
func NewStorage() *minio.Client {
	// Inisialisasi MinIO client
	client, err := minio.New(setting.StoragePathPrefix, &minio.Options{
		Creds:  credentials.NewStaticV4(setting.StorageAccessKey, setting.StorageSecretKey, ""),
		Secure: false, // Ubah menjadi true jika menggunakan HTTPS
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO client: %v", err)
	}

	bucketName := setting.StorageBucket

	// Pastikan bucket ada
	err = ensureBucketExists(client, bucketName)
	if err != nil {
		log.Fatalf("Error ensuring bucket exists: %v", err)
	}

	return client
}

// ensureBucketExists memastikan bucket ada atau membuatnya jika belum ada
func ensureBucketExists(client *minio.Client, bucketName string) error {
	// Periksa apakah bucket sudah ada
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}

	if exists {
		log.Printf("Bucket '%s' sudah ada", bucketName)
		return nil
	}

	// Buat bucket jika belum ada
	err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}

	log.Printf("Successfully created bucket '%s'", bucketName)
	return nil
}