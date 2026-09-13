package storage

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
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

func resolvePurposeLabelEN(purpose PurposeType) string {
	switch purpose {
	case PurposeStallVerification:
		return "STALL VERIFICATION"
	case PurposeSupplierVerification:
		return "SUPPLIER VERIFICATION"
	case PurposeLeaseContract:
		return "LEASE CONTRACT"
	default:
		return "IDENTITY VERIFICATION"
	}
}

// Helper untuk mengubah gambar logo nama menjadi SILUET MONOKROM PUTIH (Opacity 35%)
func applyWhiteSilhouetteWithOpacity(src image.Image, opacity float64) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := src.At(x, y).RGBA()
			if a > 0 {
				newA := uint8((float64(a>>8) * opacity))
				dst.Set(x, y, color.NRGBA{
					R: 255,
					G: 255,
					B: 255,
					A: newA,
				})
			}
		}
	}
	return dst
}

func ApplyWatermarkFromBytes(imageBytes []byte, purpose PurposeType) ([]byte, error) {
	im, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	optimized, err := OptimizeImageBytes(imageBytes, imageProfileForFolder("/legals"))
	if err == nil {
		if optIm, _, errOpt := image.Decode(bytes.NewReader(optimized)); errOpt == nil {
			im = optIm
		}
	}

	bounds := im.Bounds()
	w := float64(bounds.Dx())
	hOriginal := float64(bounds.Dy())

	// Hitung tinggi footer banner (12% dari tinggi asli)
	footerHeight := maxFloat(hOriginal*0.12, 64)

	// EXTEND CANVAS: Tinggi total canvas baru = tinggi foto asli + tinggi footer
	hTotal := hOriginal + footerHeight

	dc := gg.NewContext(int(w), int(hTotal))

	// Render foto asli tanpa terpotong di koordinat Y = 0
	dc.DrawImage(im, 0, 0)

	workDir := mustWorkingDirectory()
	fontPath := filepath.Join(workDir, "assets", "fonts", "Inter-Bold.ttf")

	// -------------------------------------------------------------------------
	// 1. SILUET LOGO NAMA DI TENGAH FOTO ASLI (OPACITY 35% MONOKROM PUTIH)
	// -------------------------------------------------------------------------
	logoNamePath := filepath.Join(workDir, "assets", "ic_logo_name.png")
	if logoFile, err := os.Open(logoNamePath); err == nil {
		if logoImg, _, err := image.Decode(logoFile); err == nil {
			silhouetteLogo := applyWhiteSilhouetteWithOpacity(logoImg, 0.35)

			targetW := w * 0.70
			scale := targetW / float64(logoImg.Bounds().Dx())

			dc.Push()
			// Rotasi dan posisikan di tengah area foto asli (hOriginal/2)
			dc.RotateAbout(gg.Radians(-25), w/2, hOriginal/2)
			dc.ScaleAbout(scale, scale, w/2, hOriginal/2)
			dc.DrawImageAnchored(silhouetteLogo, int(w/2), int(hOriginal/2), 0.5, 0.5)
			dc.Pop()
		}
		_ = logoFile.Close()
	}

	// -------------------------------------------------------------------------
	// 2. EXTENDED FOOTER BANNER (BERADA DI LUAR/DI BAWAH FOTO ASLI)
	// -------------------------------------------------------------------------
	footerY := hOriginal // Dimulai tepat dari batas bawah foto asli

	// Background Dark Slate Solid (#0F172A)
	dc.SetRGBA255(15, 23, 42, 255)
	dc.DrawRectangle(0, footerY, w, footerHeight)
	dc.Fill()

	// Accent Line Gradien Brand (#0657BD -> #0D9488 -> #2BE576)
	grad := gg.NewLinearGradient(0, 0, w, 0)
	grad.AddColorStop(0.0, color.RGBA{R: 0x06, G: 0x57, B: 0xBD, A: 0xFF})
	grad.AddColorStop(0.55, color.RGBA{R: 0x0D, G: 0x94, B: 0x88, A: 0xFF})
	grad.AddColorStop(1.0, color.RGBA{R: 0x2B, G: 0xE5, B: 0x76, A: 0xFF})

	dc.SetStrokeStyle(grad)
	dc.SetLineWidth(maxFloat(hOriginal*0.006, 3))
	dc.DrawLine(0, footerY, w, footerY)
	dc.Stroke()

	// Font Sizing
	mainFontSize := maxFloat(footerHeight*0.24, 13)
	subFontSize := mainFontSize * 0.85

	leftMargin := w * 0.03

	// Render Logo Icon Berwarna (ic_logo.png) di Kiri Footer
	iconPath := filepath.Join(workDir, "assets", "ic_logo.png")
	if iconFile, err := os.Open(iconPath); err == nil {
		if iconImg, _, err := image.Decode(iconFile); err == nil {
			iconH := footerHeight * 0.40
			iconW := iconH * (float64(iconImg.Bounds().Dx()) / float64(iconImg.Bounds().Dy()))

			iconCtx := gg.NewContext(int(iconW), int(iconH))
			iconCtx.Scale(iconW/float64(iconImg.Bounds().Dx()), iconH/float64(iconImg.Bounds().Dy()))
			iconCtx.DrawImage(iconImg, 0, 0)

			dc.DrawImageAnchored(iconCtx.Image(), int(leftMargin+(iconW/2)), int(footerY+(footerHeight*0.30)), 0.5, 0.5)
			leftMargin += iconW + (w * 0.015)
		}
		_ = iconFile.Close()
	}

	if _, statErr := os.Stat(fontPath); statErr == nil {
		_ = dc.LoadFontFace(fontPath, mainFontSize)
	} else {
		dc.SetFontFace(basicfont.Face7x13)
	}

	// Baris 1 (Kiri): Purpose Dokumen
	dc.SetRGBA255(45, 212, 191, 255)
	purposeText := fmt.Sprintf("PURPOSE: FOR %s ONLY", resolvePurposeLabelEN(purpose))
	dc.DrawStringAnchored(purposeText, leftMargin, footerY+(footerHeight*0.30), 0, 0.5)

	// Baris 1 (Kanan): Timestamp Watermark (Format: 02-01-2006 15:04:05 MST)
	dc.SetRGBA255(255, 255, 255, 255)
	wibLocation := time.FixedZone("WIB", 7*3600)
	timestampText := fmt.Sprintf("WATERMARKED: %s", time.Now().In(wibLocation).Format("02-01-2006 15:04:05 MST"))
	dc.DrawStringAnchored(timestampText, w*0.97, footerY+(footerHeight*0.30), 1, 0.5)

	// Baris 2: Legal Warning & Safety Notice
	if _, statErr := os.Stat(fontPath); statErr == nil {
		_ = dc.LoadFontFace(fontPath, subFontSize)
	}

	dc.SetRGBA255(253, 224, 71, 255)
	warningNotice := "NOT VALID FOR OTHER TRANSACTIONS & PROHIBITED FROM MISUSE. PROTECTED BY LAPAKITA PLATFORM."
	dc.DrawStringAnchored(warningNotice, w*0.03, footerY+(footerHeight*0.72), 0, 0.5)

	var output bytes.Buffer
	if err := jpeg.Encode(&output, dc.Image(), &jpeg.Options{Quality: 92}); err != nil {
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
