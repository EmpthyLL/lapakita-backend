package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"lapakita-backend/config"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

type ImageKitService struct {
	client imagekit.Client
}

func NewImageKitService(cfg *config.Config) *ImageKitService {
	client := imagekit.NewClient(
		option.WithPrivateKey(cfg.ImageKitPrivateKey),
	)

	return &ImageKitService{
		client: client,
	}
}

func (s *ImageKitService) UploadFile(
	ctx context.Context,
	fileHeader *multipart.FileHeader,
	folder string,
) (string, error) {

	file, err := fileHeader.Open()
	if err != nil {
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
		return "", fmt.Errorf("upload imagekit: %w", err)
	}

	return resp.URL, nil
}
