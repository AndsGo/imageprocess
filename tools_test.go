package imageprocess

import (
	"fmt"
	"testing"
)

// parseHexColor converts a hex color string to an image/color.RGBA.
func Test_ParseHexColor(t *testing.T) {
	color, err := ParseHexColor("#1366ec")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(color)
}

// downloadImage downloads an image from a URL and returns it as an image.Image.
func Test_DownloadImage(t *testing.T) {
	img, fileName, err := DownloadImage("https://oss-console-img-demo-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/example.jpg")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(img, fileName)
}
