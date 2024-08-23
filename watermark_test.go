package imageprocess

import (
	"fmt"
	"image/color"
	"os"
	"testing"
)

func TestWarterMarkText(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	wp := TextWatermarkOption{WatermarkOption: WatermarkOption{
		Opacity: 100, Position: Center, X: 10, Y: 10,
	}, Color: &color.RGBA{111, 11, 22, 1}, Size: 120, Text: "hello watermark", Rotate: 90}
	img = WarterMarkText(img, wp)
	file, _ := os.Create(fmt.Sprintf("examples/out.%s", f))
	EncodeImage(img, file, PNG, 100)
}
