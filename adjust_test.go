package imageprocess

import (
	"testing"
)

func Test_AdjustBlur(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	img = AdjustBlur(img, BlurOption{Radius: 10})
	SaveImage(img, "examples/out.jpg", f, 100)
}

func Test_AddjustSharpen(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	img = AddjustSharpen(img, SharpenOption{Value: 1000})
	SaveImage(img, "examples/out.jpg", f, 100)
}

func Test_AddjustBright(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	img = AddjustBright(img, BrightnessOption{Value: 50})
	SaveImage(img, "examples/out.jpg", f, 100)
}

func Test_AddjustContrast(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	img = AddjustContrast(img, ContrastOption{Value: -10})
	SaveImage(img, "examples/out.jpg", f, 100)
}
