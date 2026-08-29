package storage

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"lapakita-backend/config"
	"lapakita-backend/pkg/logger"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"go.uber.org/zap"
)

type ImageKitService struct {
	client  imagekit.Client
	appName string
	log     *logger.Logger
}

func NewImageKitService(cfg *config.Config, log *logger.Logger) *ImageKitService {
	client := imagekit.NewClient(
		option.WithPrivateKey(cfg.ImageKitPrivateKey),
	)

	return &ImageKitService{
		client:  client,
		appName: strings.ToLower(strings.TrimSpace(cfg.AppName)),
		log:     log,
	}
}

// helperBuildFolderPath menyusun path folder dengan prefix nama aplikasi dari config
func (s *ImageKitService) helperBuildFolderPath(folder string) string {
	cleanFolder := strings.TrimPrefix(folder, "/")
	if s.appName == "" {
		return "/" + cleanFolder
	}
	if cleanFolder == "" {
		return "/" + s.appName
	}
	return fmt.Sprintf("/%s/%s", s.appName, cleanFolder)
}

func (s *ImageKitService) UploadFile(
	ctx context.Context,
	fileHeader *multipart.FileHeader,
	folder string,
) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		s.log.Error("[ImageKit] Failed to open uploaded file", zap.Error(err))
		return "", fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	targetFolder := s.helperBuildFolderPath(folder)

	resp, err := s.client.Files.Upload(
		ctx,
		imagekit.FileUploadParams{
			File:     file,
			FileName: fileHeader.Filename,
			Folder:   imagekit.String(targetFolder),
		},
	)
	if err != nil {
		s.log.Error("[ImageKit] Failed to upload file",
			zap.String("filename", fileHeader.Filename),
			zap.String("folder", targetFolder),
			zap.Error(err),
		)
		return "", fmt.Errorf("upload imagekit: %w", err)
	}

	return resp.URL, nil
}

// UploadFromURL memproses input gambar (Base64 atau HTTP URL) dan mengunggahnya dengan prefix folder nama aplikasi
func (s *ImageKitService) UploadFromURL(
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
			s.log.Error("[ImageKit] Failed to decode base64 string", zap.Error(err))
			return "", fmt.Errorf("decode base64: %w", err)
		}

		fileReader = bytes.NewReader(decodedBytes)
	} else {
		// 2. Jika input berupa URL HTTP/HTTPS (URL Avatar Google)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageSource, nil)
		if err != nil {
			s.log.Error("[ImageKit] Failed to create HTTP request", zap.Error(err))
			return "", fmt.Errorf("create http request: %w", err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		httpResp, err := http.DefaultClient.Do(req)
		if err != nil {
			s.log.Error("[ImageKit] Failed to download image from URL",
				zap.String("url", imageSource),
				zap.Error(err),
			)
			return "", fmt.Errorf("download image from url: %w", err)
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			s.log.Error("[ImageKit] Download image failed with non-200 status",
				zap.String("url", imageSource),
				zap.Int("status_code", httpResp.StatusCode),
			)
			return "", fmt.Errorf("download image failed with status: %s", httpResp.Status)
		}

		fileReader = httpResp.Body
	}

	targetFolder := s.helperBuildFolderPath(folder)

	// 3. Eksekusi Upload ke ImageKit menggunakan io.Reader
	resp, err := s.client.Files.Upload(
		ctx,
		imagekit.FileUploadParams{
			File:     fileReader,
			FileName: fileName,
			Folder:   imagekit.String(targetFolder),
		},
	)
	if err != nil {
		s.log.Error("[ImageKit] Failed to upload stream to ImageKit",
			zap.String("filename", fileName),
			zap.String("folder", targetFolder),
			zap.Error(err),
		)
		return "", fmt.Errorf("upload to imagekit: %w", err)
	}

	s.log.Info("[ImageKit] Avatar uploaded successfully",
		zap.String("url", resp.URL),
		zap.String("folder", targetFolder),
	)
	return resp.URL, nil
}
