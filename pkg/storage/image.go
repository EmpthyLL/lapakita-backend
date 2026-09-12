package storage

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"strings"

	"golang.org/x/image/draw"
)

type ImageProfile struct {
	MaxDimension int
	Quality      int
}

func imageProfileForFolder(folder string) ImageProfile {
	switch {
	case strings.Contains(folder, "/avatars"):
		return ImageProfile{MaxDimension: 512, Quality: 85}
	case strings.Contains(folder, "/documents") || strings.Contains(folder, "/legals"):
		return ImageProfile{MaxDimension: 2400, Quality: 90}
	default:
		return ImageProfile{MaxDimension: 1600, Quality: 85}
	}
}

func OptimizeImageBytes(input []byte, profile ImageProfile) ([]byte, error) {
	imageData, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}
	imageData = resizeImage(imageData, profile.MaxDimension)

	var output bytes.Buffer
	if err := jpeg.Encode(&output, imageData, &jpeg.Options{Quality: profile.Quality}); err != nil {
		return nil, fmt.Errorf("encode jpeg: %w", err)
	}
	return output.Bytes(), nil
}

func resizeImage(source image.Image, maxDimension int) image.Image {
	width := source.Bounds().Dx()
	height := source.Bounds().Dy()
	if maxDimension <= 0 || (width <= maxDimension && height <= maxDimension) {
		return source
	}

	scale := float64(maxDimension) / float64(width)
	if height > width {
		scale = float64(maxDimension) / float64(height)
	}
	newWidth := int(float64(width) * scale)
	newHeight := int(float64(height) * scale)
	resized := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.CatmullRom.Scale(resized, resized.Bounds(), source, source.Bounds(), draw.Over, nil)
	return resized
}

func optimizeReader(reader io.Reader, profile ImageProfile) ([]byte, error) {
	input, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read image: %w", err)
	}
	optimized, err := OptimizeImageBytes(input, profile)
	if err != nil {
		return input, nil
	}
	return optimized, nil
}
