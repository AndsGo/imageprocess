package imageprocess

import (
	"fmt"
	"image/color"
	"testing"
)

func Test_ParseOptions(t *testing.T) {
	options, err := ParseOptions("image/watermark,t_20,g_center,x_10,y_10,text_hello watermark,color_1366ec,size_200")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(options)
}

func Test_SerializeOptions(t *testing.T) {
	optionMap := make([]Option, 0)
	// optionMap = append(optionMap, Option{Blur, BlurOption{Radius: 10}})
	optionMap = append(optionMap, Option{Resize, ResizeOption{ResizeMode: Pad, Width: 100, Height: 100, Color: &color.RGBA{R: 255, G: 255, B: 0, A: 255}}})
	optionMap = append(optionMap, Option{Watermark, TextWatermarkOption{WatermarkOption: WatermarkOption{
		Opacity: 20, Position: Center, X: 10, Y: 10,
	}, Color: &color.RGBA{0, 0, 0, 1}, Size: 200, Text: "hello watermark"}})
	res := SerializeOptions(optionMap)
	t.Logf("%s\n", res)
}
