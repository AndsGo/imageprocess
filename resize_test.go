package imageprocess

import (
	"fmt"
	"image/color"
	"image/gif"
	"os"
	"testing"
)

// parseHexColor converts a hex color string to an image/color.RGBA.
func Test_ResizeImage(t *testing.T) {
	// D:\work\notes\demo\example\image\go_demo\imaging\example.gif
	img, f, err := LoadImage("examples/example.gif")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
	op := ResizeOption{ResizeMode: Pad, Width: 10, Height: 100, Limit: 1, Color: &color.RGBA{R: 255, G: 255, B: 0, A: 255}}
	img = ResizeImage(img, op)
	file, _ := os.Create("examples/image_resize.gif")
	EncodeImage(img, file, f, 100)
}

func Test_ResizeGif(t *testing.T) {
	// D:\work\notes\demo\example\image\go_demo\imaging\example.gif
	img, err := LoadGif("examples/example.gif")
	if err != nil {
		t.Error(err)
	}
	op := ResizeOption{ResizeMode: Pad, Width: 20, Height: 100, Limit: 1, Color: &color.RGBA{R: 255, G: 255, B: 0, A: 255}}
	ResizeGif(img, op)
	file, _ := os.Create("examples/image_resize.gif")
	gif.EncodeAll(file, img)
}
