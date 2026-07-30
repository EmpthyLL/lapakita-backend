package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"lapakita-backend/config"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService struct {
	client     *minio.Client
	bucketName string
	endpoint   string
	useSSL     bool
}

func NewMinioService(cfg *config.Config) (*MinioService, error) {
	client, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: cfg.MinioUseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("gagal inisialisasi MinIO client: %w", err)
	}

	service := &MinioService{
		client:     client,
		bucketName: cfg.MinioBucketName,
		endpoint:   cfg.MinioEndpoint,
		useSSL:     cfg.MinioUseSSL,
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, cfg.MinioBucketName)
	if err != nil {
		return nil, fmt.Errorf("gagal mengecek ketersediaan bucket MinIO: %w", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, cfg.MinioBucketName, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("gagal membuat bucket MinIO: %w", err)
		}
	}

	return service, nil
}

func (s *MinioService) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader, folderPath string) (string, string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", "", fmt.Errorf("gagal membuka file: %w", err)
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	objectName := fmt.Sprintf("%s/%s%s", folderPath, uuid.New().String(), ext)
	contentType := fileHeader.Header.Get("Content-Type")

	_, err = s.client.PutObject(ctx, s.bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", "", fmt.Errorf("gagal upload file ke MinIO: %w", err)
	}

	protocol := "http"
	if s.useSSL {
		protocol = "https"
	}
	fileURL := fmt.Sprintf("%s://%s/%s/%s", protocol, s.endpoint, s.bucketName, objectName)

	return objectName, fileURL, nil
}

func (s *MinioService) GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	presignedURL, err := s.client.PresignedGetObject(ctx, s.bucketName, objectName, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("gagal generate presigned URL: %w", err)
	}
	return presignedURL.String(), nil
}

func (s *MinioService) DeleteFile(ctx context.Context, objectName string) error {
	err := s.client.RemoveObject(ctx, s.bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("gagal menghapus file dari MinIO: %w", err)
	}
	return nil
}
