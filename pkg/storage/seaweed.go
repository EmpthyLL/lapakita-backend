package storage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"strings"

	"lapakita-backend/config"
	"lapakita-backend/pkg/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

type SeaweedFSService struct {
	client     *s3.Client
	bucketName string
	endpoint   string
	appName    string
	useSSL     bool
	log        *logger.Logger
}

func NewSeaweedFSService(cfg *config.Config, log *logger.Logger) (*SeaweedFSService, error) {
	scheme := "http"
	if cfg.SeaweedFSUseSSL {
		scheme = "https"
	}
	customEndpoint := fmt.Sprintf("%s://%s", scheme, cfg.SeaweedFSEndpoint)

	s3Cfg, err := s3Config.LoadDefaultConfig(context.Background(),
		s3Config.WithRegion("us-east-1"), // Region dummy wajib untuk S3
		s3Config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.SeaweedFSAccessKey,
			cfg.SeaweedFSSecretKey,
			"",
		)),
	)
	if err != nil {
		log.Error("[SeaweedFS] Failed to load S3 configuration", zap.Error(err))
		return nil, fmt.Errorf("seaweedfs config: %w", err)
	}

	client := s3.NewFromConfig(s3Cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(customEndpoint)
		o.UsePathStyle = true // Wajib untuk S3 lokal
	})

	svc := &SeaweedFSService{
		client:     client,
		bucketName: cfg.SeaweedFSBucketName,
		endpoint:   cfg.SeaweedFSEndpoint,
		appName:    strings.ToLower(strings.TrimSpace(cfg.AppName)),
		useSSL:     cfg.SeaweedFSUseSSL,
		log:        log,
	}

	// Buat bucket secara otomatis jika belum ada di SeaweedFS
	ctx := context.Background()
	_, err = client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(cfg.SeaweedFSBucketName)})
	if err != nil {
		_, createErr := client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: aws.String(cfg.SeaweedFSBucketName)})
		if createErr != nil {
			log.Warn("[SeaweedFS] Bucket creation check failed", zap.Error(createErr))
		} else {
			log.Info("[SeaweedFS] Bucket created successfully", zap.String("bucket", cfg.SeaweedFSBucketName))
		}
	}

	return svc, nil
}

// helperBuildObjectKey menyusun path folder dan file key dengan prefix nama aplikasi dari config
func (s *SeaweedFSService) helperBuildObjectKey(folder string, fileName string) string {
	cleanFolder := strings.TrimPrefix(folder, "/")
	if s.appName != "" {
		if cleanFolder == "" {
			cleanFolder = s.appName
		} else {
			cleanFolder = path.Join(s.appName, cleanFolder)
		}
	}
	return path.Join(cleanFolder, fileName)
}

func (s *SeaweedFSService) UploadFile(
	ctx context.Context,
	fileHeader *multipart.FileHeader,
	folder string,
) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		s.log.Error("[SeaweedFS] Failed to open uploaded file", zap.Error(err))
		return "", fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	optimizedBytes, err := optimizeReader(file, imageProfileForFolder(folder))
	if err != nil {
		return "", fmt.Errorf("optimize image: %w", err)
	}

	objectKey := s.helperBuildObjectKey(folder, fileHeader.Filename)
	contentType := http.DetectContentType(optimizedBytes)

	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(optimizedBytes),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		s.log.Error("[SeaweedFS] Failed to upload file",
			zap.String("filename", fileHeader.Filename),
			zap.String("key", objectKey),
			zap.Error(err),
		)
		return "", fmt.Errorf("upload seaweedfs: %w", err)
	}

	scheme := "http"
	if s.useSSL {
		scheme = "https"
	}
	fileURL := fmt.Sprintf("%s://%s/%s/%s", scheme, s.endpoint, s.bucketName, objectKey)

	s.log.Info("[SeaweedFS] File uploaded successfully",
		zap.String("url", fileURL),
		zap.String("key", objectKey),
	)

	return fileURL, nil
}

// UploadFromURL memproses input gambar (Base64 atau HTTP URL) dan mengunggahnya ke SeaweedFS
func (s *SeaweedFSService) UploadFromURL(
	ctx context.Context,
	imageSource string,
	fileName string,
	folder string,
) (string, error) {
	var fileReader io.Reader

	// 1. Jika input berupa string Base64
	if strings.HasPrefix(imageSource, "data:image") || (!strings.HasPrefix(imageSource, "http://") && !strings.HasPrefix(imageSource, "https://")) {
		rawBase64 := imageSource
		if idx := strings.Index(imageSource, ","); idx != -1 {
			rawBase64 = imageSource[idx+1:]
		}

		decodedBytes, err := base64.StdEncoding.DecodeString(rawBase64)
		if err != nil {
			s.log.Error("[SeaweedFS] Failed to decode base64 string", zap.Error(err))
			return "", fmt.Errorf("decode base64: %w", err)
		}

		fileReader = bytes.NewReader(decodedBytes)
	} else {
		// 2. Jika input berupa URL HTTP/HTTPS (e.g. Avatar Google)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageSource, nil)
		if err != nil {
			s.log.Error("[SeaweedFS] Failed to create HTTP request", zap.Error(err))
			return "", fmt.Errorf("create http request: %w", err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		httpResp, err := http.DefaultClient.Do(req)
		if err != nil {
			s.log.Error("[SeaweedFS] Failed to download image from URL",
				zap.String("url", imageSource),
				zap.Error(err),
			)
			return "", fmt.Errorf("download image from url: %w", err)
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			s.log.Error("[SeaweedFS] Download image failed with non-200 status",
				zap.String("url", imageSource),
				zap.Int("status_code", httpResp.StatusCode),
			)
			return "", fmt.Errorf("download image failed with status: %s", httpResp.Status)
		}

		fileReader = httpResp.Body
	}

	optimizedBytes, err := optimizeReader(fileReader, imageProfileForFolder(folder))
	if err != nil {
		return "", fmt.Errorf("optimize image: %w", err)
	}

	objectKey := s.helperBuildObjectKey(folder, fileName)
	contentType := http.DetectContentType(optimizedBytes)

	// 3. Eksekusi Upload ke SeaweedFS
	_, err = s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(optimizedBytes),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		s.log.Error("[SeaweedFS] Failed to upload stream to SeaweedFS",
			zap.String("filename", fileName),
			zap.String("key", objectKey),
			zap.Error(err),
		)
		return "", fmt.Errorf("upload to seaweedfs: %w", err)
	}

	scheme := "http"
	if s.useSSL {
		scheme = "https"
	}
	fileURL := fmt.Sprintf("%s://%s/%s/%s", scheme, s.endpoint, s.bucketName, objectKey)

	s.log.Info("[SeaweedFS] Image uploaded successfully from URL/Base64",
		zap.String("url", fileURL),
		zap.String("key", objectKey),
	)

	return fileURL, nil
}
