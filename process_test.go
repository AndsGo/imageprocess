package imageprocess

import (
	"image/color"
	"os"
	"testing"
)

func Test_Process(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	file, _ := os.Create("examples/out.jpg")
	options := make([]Option, 0)
	// 格式转换为PNG
	options = append(options, Option{FormatType, FormatOption{Format: PNG}})
	//  改变大小
	options = append(options, Option{Resize, ResizeOption{ResizeMode: Pad, Width: 300, Height: 300, Color: &color.RGBA{R: 255, G: 255, B: 0, A: 255}}})
	// 裁剪
	options = append(options, Option{Crop, CropOption{Width: 200, Height: 200, X: 0, Y: 0, Position: Center}})
	// 水印
	options = append(options, Option{Watermark, TextWatermarkOption{WatermarkOption: WatermarkOption{
		Opacity: 100, Position: South, X: 10, Y: 10,
	}, Color: &color.RGBA{111, 222, 111, 1}, Size: 40, Text: "hello watermark"}})
	// 模糊
	options = append(options, Option{Blur, GammaOption{Value: 5}})
	// 质量
	options = append(options, Option{Quality, QualityOption{Quality: 500}})
	err = Process(img, file, f, options)
	if err != nil {
		t.Error(err)
	}
}

func Test_ProcessGif(t *testing.T) {
	img, err := LoadGif("examples/example.gif")
	if err != nil {
		t.Error(err)
	}
	file, _ := os.Create("examples/out.gif")
	options := make([]Option, 0)
	// 格式转换为GIF
	options = append(options, Option{FormatType, FormatOption{Format: GIF}})
	options = append(options, Option{Gamma, GammaOption{Value: 500}})
	options = append(options, Option{Resize, ResizeOption{ResizeMode: Pad, Width: 300, Height: 300, Color: &color.RGBA{R: 255, G: 255, B: 255, A: 1}}})
	options = append(options, Option{Crop, CropOption{Width: 200, Height: 200, X: 0, Y: 0, Position: Center}})
	options = append(options, Option{Watermark, TextWatermarkOption{WatermarkOption: WatermarkOption{
		Opacity: 20, Position: Center, X: 10, Y: 10,
	}, Color: &color.RGBA{0, 0, 0, 1}, Size: 40, Text: "hello watermark"}})
	err = ProcessGif(img, file, options)
	if err != nil {
		t.Error(err)
	}
}

func Test_UrlOptions(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	file, _ := os.Create("examples/out.jpg")
	// 增加水印，然后修改大小
	options, err := ParseOptions("image/watermark,t_30,g_center,x_10,y_10,text_hello watermark,color_1366ec,size_200/resize,m_pad,h_100,w_100,color_FF0000")
	if err != nil {
		t.Error(err)
	}
	err = Process(img, file, f, options)
	if err != nil {
		t.Error(err)
	}

}
