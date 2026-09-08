package storage

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"

	"github.com/fogleman/gg"
)

type PurposeType string

const (
	PurposeStallVerification    PurposeType = "stall_verification"
	PurposeSupplierVerification PurposeType = "supplier_verification"
	PurposeLeaseContract        PurposeType = "lease_contract"
)

func ApplyWatermarkFromBytes(imageBytes []byte, purpose PurposeType) ([]byte, error) {
	im, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	bounds := im.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	dc := gg.NewContext(w, h)
	dc.DrawImage(im, 0, 0)

	label := "FOR LAPAKITA VERIFICATION ONLY"
	switch purpose {
	case PurposeStallVerification:
		label = "STALL VERIFICATION DOCUMENT - LAPAKITA"
	case PurposeSupplierVerification:
		label = "SUPPLIER VERIFICATION DOCUMENT - LAPAKITA"
	case PurposeLeaseContract:
		label = "LEASE CONTRACT DOCUMENT - LAPAKITA"
	}

	// Brand Gradient (#0657BD -> #0D9488 -> #2BE576)
	grad := gg.NewLinearGradient(0, 0, float64(w), float64(h))
	grad.AddColorStop(0.0, color.RGBA{R: 0x06, G: 0x57, B: 0xBD, A: 0xFF})
	grad.AddColorStop(0.55, color.RGBA{R: 0x0D, G: 0x94, B: 0x88, A: 0xFF})
	grad.AddColorStop(1.0, color.RGBA{R: 0x2B, G: 0xE5, B: 0x76, A: 0xFF})

	dc.Push()
	dc.RotateAbout(gg.Radians(-30), float64(w)/2, float64(h)/2)

	fontSize := float64(w) / 14.0
	if err := dc.LoadFontFace("assets/fonts/Inter-Bold.ttf", fontSize); err != nil {
		dc.SetRGB(1, 1, 1)
	}

	dc.SetStrokeStyle(grad)
	dc.SetFillStyle(grad)

	dc.SetRGBA255(6, 87, 189, 140)
	dc.DrawStringAnchored(label, float64(w)/2, float64(h)/2, 0.5, 0.5)
	dc.DrawStringAnchored("NOT VALID FOR OTHER TRANSACTIONS", float64(w)/2, (float64(h)/2)+fontSize*1.2, 0.5, 0.5)

	dc.Pop()

	var buf bytes.Buffer
	if err := dc.EncodePNG(&buf); err != nil {
		return nil, fmt.Errorf("encode png: %w", err)
	}

	return buf.Bytes(), nil
}
