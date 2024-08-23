**English** | [中文](./README.md) 
# imgaeprocess

This project is based on the following open-source libraries:
[imaging]([disintegration/imaging: Imaging is a simple image processing package for Go (github.com)](https://github.com/disintegration/imaging)) About Imaging is a simple image processing package for Go

[webp](https://github.com/chai2010/webp) WebP decoder and encoder for Go (Zero Dependencies).

[gg](https://github.com/fogleman/gg) Go Graphics - 2D rendering in Go with a simple API.

`Imageprocess`is a simple Go image processing package that supports WEBP, JPG, JPEG, PNG, BMP, TIFF, and GIF. It provides image processing capabilities similar to those of Alibaba Cloud's OSS, suitable for building a local file OSS image processing system.
[Resize images](doc/resize_en.md)

[Add watermarks](doc/watermark_en.md)

[Custom crop](doc/crop_en.md)

[Adjust image quality](doc/quality_en.md)

[Convert image formats](doc/format_en.md)

[Blur](doc/blur_en.md)

[Rotate](doc/rotate_en.md)

[Brightness](doc/bright_en.md)

[Sharpen](doc/sharpen_en.md)

[Contrast](doc/sharpen_en.md)

## 0.**Installation**

```shell
go get github.com/AndsGo/imageprocess
```

## 1.**Code Example**

Here is an example of how to use the API in code:

```go
// Resize the image to a height of 100px and a width of 300px, with the mode being proportional scaling to match the largest side
img = imageprocess.ResizeImage(img, ResizeOption{ResizeMode: Lfit, Width: 10, Height: 100)
                                                 
// Scale the original image to a width and height of 100px: `resize,h_100,w_100` scaling mode fill: `m_fill`
img = imageprocess.ResizeImage(img, ResizeOption{ResizeMode: Fill, Width: 10, Height: 100)
                                                 
// Scale the original image to a width and height of 100px: `resize,h_100,w_100` scaling mode pad: `m_pad`. Fill with red: `color_FF0000`
img = imageprocess.ResizeImage(img, ResizeOption{ResizeMode: Pad, Width: 10, Height: 100 Color: &color.RGBA{R: 255, G: 0, B: 0, A: 255}})
                                                 
// Thumbnail example.jpg to a width and height of 300: `resize,w_300,h_300` watermark content "Hello World": `text_Hello%20World` watermark text color is white, font size is 30: `color_FFFFFF,size_30` watermark text position is bottom right, horizontal margin 10, vertical offset from the center 10: `g_se,x_10,y_10`
// Multiple mode serial processing can be done using the Process or (ProcessGif) method, and you can also use all the Process or (ProcessGif) methods for processing
options := make([]Option, 0)
options = append(options, Option{Resize, ResizeOption{Width: 300, Height: 300)
options = append(options, Option{Watermark, TextWatermarkOption{WatermarkOption: WatermarkOption{
    Opacity: 20, Position: Center, X: 10, Y: 10,
}, Color: &color.RGBA{255, 255, 255, 1}, Size: 30, Text: "Hello World"}})
imageprocess.Process(img, file, options)
                                                      
// Crop the image starting at coordinates (800,500) with a crop area of 300 pixels by 300 pixels.
img = imageprocess.CropImage(img, CropOption{X: 800, Y: 500, Width: 300, Height: 300})

// Crop the image starting at the bottom right corner of the original image with a crop area of 900 pixels by 900 pixels.
img = imageprocess.CropImage(img, CropOption{Position: SouthEast, Width: 900, Height: 900})

// Resize the original image to a width of 100 pixels and set the image quality to 80%.
options := make([]Option, 0)
options = append(options, Option{Resize, ResizeOption{Width: 100, Height: 100}})

// Append quality adjustment to the options with a quality setting of 80%.
options = append(options, Option{Quality, QualityOption{Quality: 80}})

// Convert the format of the image to PNG and append this to the options.
options = append(options, Option{FormatType, FormatOption{Format: PNG}})

// Process the image with the specified options.
imageprocess.Process(img, file, options)

// Rotate the original image 90 degrees clockwise.
img = imageprocess.AddjustRotate(img, RotateOption{Value: 90})

// Increase the brightness of the image by 50 units.
img = imageprocess.AddjustBright(img, BrightnessOption{Value: 50})

// Apply sharpening to the original image with a sharpening parameter of 100.
img = imageprocess.AddjustSharpen(img, SharpenOption{Value: 100})

// Increase the contrast of the image by 50 units.
img = imageprocess.AddjustContrast(img, ContrastOption{Value: 50})  
```

For a comprehensive code example, see `process_test.go`

```go
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
	file, _ := os.Create("examples/out.PNG")
	options := make([]Option, 0)
	// Convert to PNG
	options = append(options, Option{FormatType, FormatOption{Format: PNG}})
	//  Resize
	options = append(options, Option{Resize, ResizeOption{ResizeMode: Pad, Width: 300, Height: 300, Color: &color.RGBA{R: 255, G: 255, B: 0, A: 255}}})
	// Crop
	options = append(options, Option{Crop, CropOption{Width: 200, Height: 200, X: 0, Y: 0, Position: Center}})
	// Watermark
	options = append(options, Option{Watermark, TextWatermarkOption{WatermarkOption: WatermarkOption{
		Opacity: 100, Position: South, X: 10, Y: 10,
	}, Color: &color.RGBA{111, 222, 111, 1}, Size: 40, Text: "hello watermark"}})
	// Blur
	options = append(options, Option{Blur, GammaOption{Value: 5}})
	// Quality
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
	// Convert to GIF
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
	// Add a watermark and then modify the size
	options, err := ParseOptions("image/watermark,t_30,g_center,x_10,y_10,text_hello watermark,color_1366ec,size_200/resize,m_pad,h_100,w_100,color_FF0000")
	if err != nil {
		t.Error(err)
	}
	err = Process(img, file, f, options)
	if err != nil {
		t.Error(err)
	}

}
```



| Meaning                                                      | Image                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| 改变大小，高100px，宽300px，模式等比缩放匹配最大边           | ![break](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527171.jpg) |
| 将原图缩放成宽高100 px：`resize,h_100,w_100` 缩放模式fill：`m_fill` | ![自动裁剪](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527179.jpg) |
| 将原图缩放成宽高100 px：`resize,h_100,w_100` 缩放模式pad：`m_pad`。 以红色填充：`color_FF0000` | ![填充红色](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527183.jpg) |
| 将example.jpg缩略为宽高300：`resize,w_300,h_300` 水印内容为“Hello World”：`text_Hello%20World 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30` ` 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10` | ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529186.jpg) |
| 裁剪起点为（800,500）：`crop,x_800,y_500` 裁减范围300 px*300 px：`w_300,h_300` | ![裁剪2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674612.jpg) |
| 裁剪起点为原图右下角：`crop,g_se` 裁减范围900 px*900 px：`w_900,h_900` | ![裁剪3](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674614.jpg) |
| 原图缩放为宽100 px：`resize,w_100,h_100` 图片相对质量设置为80%：`quality,q_80` | ![变换1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8442799661/p529279.jpg) |
| 将原图转换为PNG格式                                          | ![png](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139213.png) |
| 将原图按顺时针旋转90°                                        | ![旋转1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529612.jpg) |
| 将图片亮度提高50                                             | ![亮度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7532220761/p529917.jpg) |
| 对原图进行锐化处理，锐化参数为100                            | ![锐化1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529929.jpg) |
| 对比度提高50                                                 | ![对比度2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529938.jpg) |

## 2.Parameter example

URL parameter example

Format conversion can be tested in `process_test`, the test code is as follows

```go
func Test_UrlOptions(t *testing.T) {
	img, f, err := LoadImage("examples/example.jpg")
	if err != nil {
		t.Error(err)
	}
	file, _ := os.Create("examples/out.jpg")
	// Add watermark and then modify the size
	options, err := ParseOptions("image/watermark,t_30,g_center,x_10,y_10,text_hello watermark,color_1366ec,size_200/resize,m_pad,h_100,w_100,color_FF0000")
	if err != nil {
		t.Error(err)
	}
    // Prcocess image
	err = Process(img, file, f, options)
	if err != nil {
		t.Error(err)
	}
}
```

| Options                                                      | Meaning                                                      | Image                                                        |
| :----------------------------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| resize,h_100,w_300,m_lfit                                    | 改变大小，高100px，宽300px，模式等比缩放匹配最大边           | ![break](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527171.jpg) |
| resize,m_fill,h_100,w_100                                    | 将原图缩放成宽高100 px：`resize,h_100,w_100` 缩放模式fill：`m_fill` | ![自动裁剪](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527179.jpg) |
| resize,m_pad,h_100,w_100,color_FF0000                        | 将原图缩放成宽高100 px：`resize,h_100,w_100` 缩放模式pad：`m_pad`。 以红色填充：`color_FF0000` | ![填充红色](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527183.jpg) |
| resize,w_300,h_300/watermark,size_30,text_Hello%20World,color_FFFFFF,g_se,x_10,y_10 | 将example.jpg缩略为宽高300：`resize,w_300,h_300` 水印内容为“Hello World”：`text_Hello%20World 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30` ` 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10` | ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529186.jpg) |
| crop,x_800,y_500,w_300,h_300                                 | 裁剪起点为（800,500）：`crop,x_800,y_500` 裁减范围300 px*300 px：`w_300,h_300` | ![裁剪2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674612.jpg) |
| crop,w_900,h_900,g_se                                        | 裁剪起点为原图右下角：`crop,g_se` 裁减范围900 px*900 px：`w_900,h_900` | ![裁剪3](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674614.jpg) |
| resize,w_100/quality,q_80                                    | 原图缩放为宽100 px：`resize,w_100` 图片相对质量设置为80%：`quality,q_80` | ![变换1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8442799661/p529279.jpg) |
| format,png                                                   | 将原图转换为PNG格式                                          | ![png](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139213.png) |
| rotate,90                                                    | 将原图按顺时针旋转90°                                        | ![旋转1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529612.jpg) |
| bright,50                                                    | 将图片亮度提高50                                             | ![亮度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7532220761/p529917.jpg) |
| sharpen,100                                                  | 对原图进行锐化处理，锐化参数为100                            | ![锐化1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529929.jpg) |
| contrast,-50                                                 | 对比度提高50                                                 | ![对比度2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529938.jpg) |

## 3.Comprehensive example

This is an <span id="comprehensive">Comprehensive example</span>of a simple file server. The code is located in the `examples` folder

```shell
cd examples
go run example.go
```

Original image: 2500*1875 ![image](./examples/example.jpg) 

Visit:

http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_500,h_300/watermark,t_80,g_se,x_10,y_10,text_hello,color_FFFFFF,size_40/format,webp

Reult: 400*300  65.4k

![imgae](./doc/1.webp)

Conversion code means:

`resize,w_500,h_300`Convert width 500, height 300

`watermark,t_80,g_se,x_10,y_10,text_hello,color_FFFFFF,size_40`Add watermark, the watermark position is located in the lower right, the distance from the edge is 10, the watermark content is hello, the color is FFFFFF, and the text size is 40

`format,webp`Format conversion to `webp`

Sample code (you need to modify **`fileFolders`** when testing your own pictures):

```go
package main

import (
	"fmt"
	"image/gif"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/AndsGo/imageprocess"
)

// You need change it
var fileFolders = "./"

func main() {
	http.HandleFunc("/file/", fileHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the file name
	fileName := strings.TrimPrefix(r.URL.Path, "/file/")
	// Open the file
	file, err := os.Open(fmt.Sprintf("%s%s", fileFolders, fileName))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	// Get parameters
	// Get the file suffix
	f, err := imageprocess.FormatFromExtension(fileName)
	if err != nil {
		// Write the processed file content to the response
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	// Processing parameters ossParameter
	ossParams := r.URL.Query().Get("x-oss-process")
	if ossParams == "" {
		// nothing to do
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	options, err := imageprocess.ParseOptions(ossParams)
	if err != nil {
		http.Error(w, fmt.Sprintf("ParseOptions %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if len(options) == 0 {
		// nothing to do
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	// Process the image 
	err = processImg(file, w, f, options)
	if err != nil {
		http.Error(w, fmt.Sprintf("processFile %s", err.Error()), http.StatusInternalServerError)
	}
}

// Convert
func processImg(file io.Reader, w io.Writer, f imageprocess.Format, options []imageprocess.Option) error {
	if f == imageprocess.GIF {
		imgGif, err := gif.DecodeAll(file)
		if err != nil {
			return err
		}
		return imageprocess.ProcessGif(imgGif, w, options)
	} else {
		img, err := imageprocess.DecodeImage(file, f)
		if err != nil {
			return err
		}
		return imageprocess.Process(img, w, f, options)
	}
}

```
