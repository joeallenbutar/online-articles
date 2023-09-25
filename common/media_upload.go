package common

import (
	"fmt"
	"mime/multipart"
	"online-articles/configuration"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadFileToMinIO(ctx *fiber.Ctx, file *multipart.FileHeader, config configuration.Config) (string, string, error) {
	// Initialize MinIO client
	endpoint := config.Get("MINIO_ENDPOINT")
	accessKey := config.Get("MINIO_ACCESSKEY")
	secretKey := config.Get("MINIO_SECRETKEY")
	bucketName := config.Get("MINIO_BUCKET")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
		Region: "us-east-1",
	})
	if err != nil {
		return "", "", err
	}

	// Upload file to MinIO
	objectName := file.Filename
	object, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer object.Close()

	_, err = minioClient.PutObject(ctx.Context(), bucketName, objectName, object, file.Size, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", "", err
	}

	// Construct the URL for the uploaded object
	url := fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)

	return url, objectName, nil
}
