package configuration

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioBucket(config Config) error {
	endpoint := config.Get("MINIO_ENDPOINT")
	accessKey := config.Get("MINIO_ACCESSKEY")
	secretKey := config.Get("MINIO_SECRETKEY")
	useSSL := false
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
		Region: "us-east-1",
	})
	if err != nil {
		fmt.Println("Error initializing MinIO client:", err)
		return nil
	}
	bucketName := config.Get("MINIO_BUCKET")
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		fmt.Println("Error creating bucket:", err)
		return nil
	}
	return nil
}
