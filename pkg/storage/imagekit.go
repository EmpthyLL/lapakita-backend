package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"lapakita-backend/config"
	"lapakita-backend/pkg/logger"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"go.uber.org/zap"
)

type ImageKitService struct {
	client imagekit.Client
	log    *logger.Logger
}

func NewImageKitService(cfg *config.Config, log *logger.Logger) *ImageKitService {
	client := imagekit.NewClient(
		option.WithPrivateKey(cfg.ImageKitPrivateKey),
		option.WithBaseURL(cfg.ImageKitUrlEndpoint),
	)

	return &ImageKitService{
		client: client,
		log:    log,
	}
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

	resp, err := s.client.Files.Upload(
		ctx,
		imagekit.FileUploadParams{
			File:     file,
			FileName: fileHeader.Filename,
			Folder:   imagekit.String(folder),
		},
	)
	if err != nil {
		s.log.Error("[ImageKit] Failed to upload file",
			zap.String("filename", fileHeader.Filename),
			zap.Error(err),
		)
		return "", fmt.Errorf("upload imagekit: %w", err)
	}

	return resp.URL, nil
}

// UploadFromURL mendownload buffer gambar dari URL lalu mengunggah io.Reader ke ImageKit
func (s *ImageKitService) UploadFromURL(
	ctx context.Context,
	imageURL string,
	fileName string,
	folder string,
) (string, error) {
	// 1. Download gambar dari URL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageURL, nil)
	if err != nil {
		s.log.Error("[ImageKit] Failed to create HTTP request", zap.Error(err))
		return "", fmt.Errorf("create http request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.log.Error("[ImageKit] Failed to download image from external URL",
			zap.String("url", imageURL),
			zap.Error(err),
		)
		return "", fmt.Errorf("download image from url: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		s.log.Error("[ImageKit] Download image failed with non-200 status",
			zap.String("url", imageURL),
			zap.Int("status_code", httpResp.StatusCode),
		)
		return "", fmt.Errorf("download image failed with status: %s", httpResp.Status)
	}

	// 2. Upload io.Reader ke ImageKit
	resp, err := s.client.Files.Upload(
		ctx,
		imagekit.FileUploadParams{
			File:     httpResp.Body,
			FileName: fileName,
			Folder:   imagekit.String(folder),
		},
	)
	if err != nil {
		s.log.Error("[ImageKit] Failed to upload stream to ImageKit",
			zap.String("filename", fileName),
			zap.Error(err),
		)
		return "", fmt.Errorf("upload imagekit from reader: %w", err)
	}

	s.log.Info("[ImageKit] Avatar uploaded successfully", zap.String("url", resp.URL))
	return resp.URL, nil
}
