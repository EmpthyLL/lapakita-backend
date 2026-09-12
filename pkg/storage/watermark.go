package storage

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/basicfont"
)

type PurposeType string

const (
	PurposeStallVerification    PurposeType = "stall_verification"
	PurposeSupplierVerification PurposeType = "supplier_verification"
	PurposeLeaseContract        PurposeType = "lease_contract"
)

func resolvePurposeLabel(purpose PurposeType) string {
	switch purpose {
	case PurposeStallVerification:
		return "VERIFIKASI LAPAK & PEMILIK - LAPAKITA"
	case PurposeSupplierVerification:
		return "VERIFIKASI SUPPLIER B2B - LAPAKITA"
	case PurposeLeaseContract:
		return "DOKUMEN KONTRAK SEWA - LAPAKITA"
	default:
		return "VERIFIKASI IDENTITAS - LAPAKITA"
	}
}

func ApplyWatermarkFromBytes(imageBytes []byte, purpose PurposeType) ([]byte, error) {
	im, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	optimized, err := OptimizeImageBytes(imageBytes, imageProfileForFolder("/legals"))
	if err != nil {
		return nil, err
	}
	im, _, err = image.Decode(bytes.NewReader(optimized))
	if err != nil {
		return nil, fmt.Errorf("decode optimized image: %w", err)
	}

	bounds := im.Bounds()
	w := float64(bounds.Dx())
	h := float64(bounds.Dy())

	dc := gg.NewContext(int(w), int(h))
	dc.DrawImage(im, 0, 0)

	fontPath := filepath.Join(mustWorkingDirectory(), "assets", "fonts", "Inter-Bold.ttf")
	if _, statErr := os.Stat(fontPath); statErr == nil {
		_ = dc.LoadFontFace(fontPath, maxFloat(w/42, 18))
	} else {
		dc.SetFontFace(basicfont.Face7x13)
	}

	label := fmt.Sprintf("%s - %s", resolvePurposeLabel(purpose), time.Now().Format("2006-01-02"))
	dc.Push()
	dc.RotateAbout(gg.Radians(-18), w/2, h/2)
	dc.SetRGBA(0.05, 0.25, 0.45, 0.32)
	dc.DrawStringAnchored(label, w/2, h/2, 0.5, 0.5)
	dc.Pop()

	var output bytes.Buffer
	if err := jpeg.Encode(&output, dc.Image(), &jpeg.Options{Quality: 90}); err != nil {
		return nil, fmt.Errorf("encode jpeg: %w", err)
	}
	return output.Bytes(), nil
}

func mustWorkingDirectory() string {
	workDir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return workDir
}

func maxFloat(value, minimum float64) float64 {
	if value < minimum {
		return minimum
	}
	return value
}
