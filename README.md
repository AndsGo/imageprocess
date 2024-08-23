**中文** | [English](./README_en.md) 
------

[TOC]
# imgaeprocess

本项目基于以下开源库:
[imaging](https://github.com/disintegration/imaging)

About Imaging is a simple image processing package for Go

[webp](https://github.com/chai2010/webp) 

WebP decoder and encoder for Go (Zero Dependencies).

[gg](https://github.com/fogleman/gg) 

Go Graphics - 2D rendering in Go with a simple API.

`Imageprocess`是一个简单的Go图像处理包。支持 WEBP,JPG,JPEG,PNG,BMP,TIFF,GIF。提供了类似阿里云`oss`的图片处理能力包括:

[图片缩放](doc/resize.md)

[图片水印](doc/watermark.md)

[自定义裁剪](doc/crop.md)

[质量变换](doc/quality.md)

[格式转换](doc/format.md)

[模糊效果](doc/blur.md)

[旋转](doc/rotate.md)

[亮度](doc/bright.md)

[锐化](doc/sharpen.md)

[对比度](doc/sharpen.md)

参数兼容阿里云oss图片处理参数，可以用于搭建本地文件oss图片处理系统。

## 0.安装

```shell
go get github.com/AndsGo/imageprocess
```

## 1.代码示例

代码`Api`使用示例

```go
// 改变大小，高100px，宽300px，模式等比缩放匹配最大边
img = imageprocess.ResizeImage(img, ResizeOption{ResizeMode: Lfit, Width: 10, Height: 100)
// 将原图缩放成宽高100 px：`resize,h_100,w_100` 缩放模式fill：`m_fill`
img = imageprocess.ResizeImage(img, ResizeOption{ResizeMode: Fill, Width: 10, Height: 100)
// 将原图缩放成宽高100 px：`resize,h_100,w_100` 缩放模式pad：`m_pad`。 以红色填充：`color_FF0000`
img = imageprocess.ResizeImage(img, ResizeOption{ResizeMode: Pad, Width: 10, Height: 100 Color: &color.RGBA{R: 255, G: 0, B: 0, A: 255}})
// 将example.jpg缩略为宽高300：`resize,w_300,h_300` 水印内容为“Hello World”：`text_Hello%20World 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30` ` 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10`
// 多模式串行处理可以使用Process或(ProcessGif)方法，你也可以全部使用Process或(ProcessGif)方法进行处理
options := make([]Option, 0)
options = append(options, Option{Resize, ResizeOption{Width: 300, Height: 300)
options = append(options, Option{Watermark, TextWatermarkOption{WatermarkOption: WatermarkOption{
    Opacity: 20, Position: Center, X: 10, Y: 10,
}, Color: &color.RGBA{255, 255, 255, 1}, Size: 30, Text: "Hello World"}})
imageprocess.Process(img, file, options)
// 裁剪起点为（800,500）：`crop,x_800,y_500` 裁减范围300 px*300 px：`w_300,h_300`
img = imageprocess.CropImage(img, CropOption{X:800,Y:500,Width: 300, Height: 300})
// 裁剪起点为原图右下角：`crop,g_se` 裁减范围900 px*900 px：`w_900,h_900`
img = imageprocess.CropImage(img, CropOption{Position: SouthEast,Width: 900, Height: 900})
// 原图缩放为宽100 px：`resize,w_100,h_100` 图片相对质量设置为80%：`quality,q_80`
options := make([]Option, 0)
options = append(options, Option{Resize, ResizeOption{Width: 100, Height: 100)
// 质量
options = append(options, Option{Quality, QualityOption{Quality: 80}})
// 格式转换为PNG
options = append(options, Option{FormatType, FormatOption{Format: PNG}})
imageprocess.Process(img, file, options)
// 将原图按顺时针旋转90° 
img = imageprocess.AddjustRotate(img,RotateOption{Value:90})
// 将图片亮度提高50
img = imageprocess.AddjustBright(img,BrightnessOption{Value:50})
// 对原图进行锐化处理，锐化参数为100
img = imageprocess.AddjustSharpen(img,SharpenOption{Value:100})
// 对比度提高50
img = imageprocess.AddjustContrast(img,ContrastOption{Value:50})     
```

综合代码示例请查看

`process_test.go`

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

## 2.参数示例

url参数示例

格式转换可以在process_test中进行测试 ,测试代码如下

```go
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
    // 处理图片
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
| resize,w_300,h_300/watermark,size_30,text_Hello World,color_FFFFFF,g_se,x_10,y_10 | 将example.jpg缩略为宽高300：`resize,w_300,h_300` 水印内容为“Hello World”：`text_Hello%20World 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30` ` 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10` | ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529186.jpg) |
| crop,x_800,y_500,w_300,h_300                                 | 裁剪起点为（800,500）：`crop,x_800,y_500` 裁减范围300 px*300 px：`w_300,h_300` | ![裁剪2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674612.jpg) |
| crop,w_900,h_900,g_se                                        | 裁剪起点为原图右下角：`crop,g_se` 裁减范围900 px*900 px：`w_900,h_900` | ![裁剪3](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674614.jpg) |
| resize,w_100/quality,q_80                                    | 原图缩放为宽100 px：`resize,w_100` 图片相对质量设置为80%：`quality,q_80` | ![变换1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8442799661/p529279.jpg) |
| format,png                                                   | 将原图转换为PNG格式                                          | ![png](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139213.png) |
| rotate,90                                                    | 将原图按顺时针旋转90°                                        | ![旋转1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529612.jpg) |
| bright,50                                                    | 将图片亮度提高50                                             | ![亮度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7532220761/p529917.jpg) |
| sharpen,100                                                  | 对原图进行锐化处理，锐化参数为100                            | ![锐化1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529929.jpg) |
| contrast,-50                                                 | 对比度提高50                                                 | ![对比度2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529938.jpg) |

## 3.综合示例

这是一个简单文件服务器的例子，代码位于`examples`文件夹下

```shell
cd examples
go run example.go
```

原图: 2500*1875 ![image](./examples/example.jpg) 

访问:

http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_500,h_300/watermark,t_80,g_se,x_10,y_10,text_hello,color_FFFFFF,size_40/format,webp

结果: 400*300  65.4k

![imgae](./doc/1.webp)

转换代码表示:

`resize,w_500,h_300`  转换宽500,高300

`watermark,t_80,g_se,x_10,y_10,text_hello,color_FFFFFF,size_40` 增加水印,水印位置位于右下,离边缘距离为10,水印内容为hello,颜色为FFFFFF,文字大小为40

`format,webp`格式转换为 `webp`

示例代码(你测试自己的图片需要修改 **`fileFolders`**):

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

// 文件夹，you need change it
var fileFolders = "./"

func main() {
	http.HandleFunc("/file/", fileHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// 获取文件名称
	fileName := strings.TrimPrefix(r.URL.Path, "/file/")
	// 打开文件
	file, err := os.Open(fmt.Sprintf("%s%s", fileFolders, fileName))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	// 获取参数
	// 获取文件后缀
	f, err := imageprocess.FormatFromExtension(fileName)
	if err != nil {
		// 将处理后的文件内容写入响应
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	//处理处理参数
	ossParams := r.URL.Query().Get("x-oss-process")
	if ossParams == "" {
		//无需处理
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
		//无需处理
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	//处理图片
	err = processImg(file, w, f, options)
	if err != nil {
		http.Error(w, fmt.Sprintf("processFile %s", err.Error()), http.StatusInternalServerError)
	}
}

// 进行转换
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
