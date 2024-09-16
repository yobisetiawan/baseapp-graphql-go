package utils

import (
	"baseapp/app/configs"
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"net/url"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func StorageCreateThumbnail(imgData []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}

	// Set the desired width
	desiredWidth := 100

	// Calculate the aspect ratio
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()
	aspectRatio := float64(originalHeight) / float64(originalWidth)

	// Calculate the height based on the aspect ratio
	calculatedHeight := int(float64(desiredWidth) * aspectRatio)

	thumb := imaging.Thumbnail(img, desiredWidth, calculatedHeight, imaging.Lanczos)

	// Encode thumbnail to JPEG or PNG and return as bytes
	var buf bytes.Buffer
	err = jpeg.Encode(&buf, thumb, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func StorageUploadToS3(fileName string, fileData []byte) error {
	s3Endpoint := configs.AppConfig.S3Endpoint
	s3AccessKeyID := configs.AppConfig.S3AccessKeyID
	s3SecretAccessKey := configs.AppConfig.S3SecretAccessKey
	s3UseSSL := configs.AppConfig.S3UseSSL == "true"
	s3BucketName := configs.AppConfig.S3BucketName

	minioClient, err := minio.New(s3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3AccessKeyID, s3SecretAccessKey, ""),
		Secure: s3UseSSL,
	})
	if err != nil {
		return err
	}

	_, err = minioClient.PutObject(
		context.Background(),
		s3BucketName,
		fileName,
		bytes.NewReader(fileData),
		int64(len(fileData)),
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func StorageGetPresignedURL(fileName string) (string, error) {
	s3Endpoint := configs.AppConfig.S3Endpoint
	s3AccessKeyID := configs.AppConfig.S3AccessKeyID
	s3SecretAccessKey := configs.AppConfig.S3SecretAccessKey
	s3UseSSL := configs.AppConfig.S3UseSSL == "true"
	s3BucketName := configs.AppConfig.S3BucketName

	minioClient, err := minio.New(s3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s3AccessKeyID, s3SecretAccessKey, ""),
		Secure: s3UseSSL,
	})
	if err != nil {
		return "", err
	}

	// Set the expiration time for the presigned URL
	expires := time.Hour * 24 // URL valid for 24 hours

	// Generate the presigned URL
	reqParams := make(url.Values)
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), s3BucketName, fileName, expires, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func StorageIsImage(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".webp"
}
