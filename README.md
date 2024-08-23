# imgaeprocess

本项目基于以下开源库
[imaging]([disintegration/imaging: Imaging is a simple image processing package for Go (github.com)](https://github.com/disintegration/imaging))

[webp](https://github.com/chai2010/webp)

[gg](https://github.com/fogleman/gg)

`Imageprocess`是一个简单的Go图像处理包。支持 WEBP,JPG,JPEG,PNG,BMP,TIFF,GIF。提供了类似阿里云`oss`的图片处理能力包括:

[图片缩放]: #图片缩放
[图片水印]: #图片水印
[自定义裁剪]: #自定义裁剪
[质量变换]: #质量变换
[格式转换]: #格式转换
[模糊效果]: #模糊效果
[旋转]: #旋转
[亮度]: #亮度
[锐化]: #锐化
[对比度]: #对比度

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
| resize,w_300,h_300/watermark,size_30,text_Hello%20World,color_FFFFFF,g_se,x_10,y_10 | 将example.jpg缩略为宽高300：`resize,w_300,h_300` 水印内容为“Hello World”：`text_Hello%20World 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30` ` 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10` | ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529186.jpg) |
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

原图: 2500*1875 ,10M![image](./examples/example.png) 

访问:

http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,w_500,h_300/watermark,t_80,g_se,x_10,y_10,text_hello,color_FFFFFF,size_40/format,webp

结果: 400*300  65.4k

![imgae](./doc/1.webp)

转换代码表示:

`resize,w_500,h_300`  

转换宽500,高300

`watermark,t_80,g_se,x_10,y_10,text_hello,color_FFFFFF,size_40` 

增加水印,水印位置位于右下,离边缘距离为10,水印内容为hello,颜色为FFFFFF,文字大小为40

`format,webp`

格式转换为 `webp`

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

# 图片缩放

更新时间：2024-08-22 13:53:25

## 注意事项

- 原图限制

  - 图片格式只能是：JPG、PNG、BMP、GIF、WebP、TIFF。

  - 宽或高不能超过30,000 px，且总像素不能超过2.5亿 px。

    动态图片（例如GIF图片）的像素计算方式为`宽*高*图片帧数`；非动态图片（例如PNG图片）的像素计算方式为`宽*高`。

- 缩放图限制

  宽或高不能超过16,384 px，且总像素不能超过16,777,216 px。

- 缩放优先级

  如果图片处理URL中同时指定按宽高缩放和等比缩放参数，则只执行指定宽高缩放。

- 缩放时只指定宽度或者高度

  - 等比缩放时，会按比例缩放图片。例如原图为200 px*100 px，将高缩放为100 px，则宽缩放为50 px。
  - 固定宽高缩放时，会将原图宽高按照指定值进行缩放。例如原图为200 px*100 px，将高缩放为100 px，则宽也缩放为100 px。

## 参数说明

操作名称：`resize`

### 指定宽高缩放

- 参数说明

  | **名称**  | **是否必选**         | **描述**                                                | **取值范围**                                                 |
  | --------- | -------------------- | ------------------------------------------------------- | ------------------------------------------------------------ |
  | **m**     | 是                   | 指定缩放的模式。                                        | lfit（默认值）：等比缩放，缩放图限制为指定w与h的矩形内的最大图片。mfit：等比缩放，缩放图为延伸出指定w与h的矩形框外的最小图片。fill：将原图等比缩放为延伸出指定w与h的矩形框外的最小图片，然后将超出的部分进行居中裁剪。pad：将原图缩放为指定w与h的矩形内的最大图片，然后使用指定颜色居中填充空白部分。fixed：固定宽高，强制缩放。 |
  | **w**     | 否                   | 指定目标缩放图的宽度。                                  | [1,16384]                                                    |
  | **h**     | 否                   | 指定目标缩放图的高度。                                  | [1,16384]                                                    |
  | **color** | 是（仅当`m为pad`时） | 当缩放模式选择为pad（缩放填充）时，可以设置填充的颜色。 | RGB颜色值，例如：000000表示黑色，FFFFFF表示白色。默认值：FFFFFF（白色） |

- 使用示例

  原图大小为200 px*100 px，缩放参数为w=150 px，h=80 px。则不同的缩略模式，得到的缩放图如下：

  lfit

  mfit

  fill

  pad

  fixed

  - 等比缩放：要求缩放图的w/h等于原图的w/h。所以，若w=150 px，则h=75 px；若h=80 px，则w=160 px。
  - 限制在指定w与h的矩形内的最大图片：即缩放图的w*h不能大于150 px*80 px。

  通过以上条件得出缩略图大小为150 px*75 px。

  ![lfit](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/4822359951/p137017.png)

## **操作方式**

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527167.jpg)

- 等比缩放

  - 按宽高缩放

    需求及处理参数如下：

    - 图片缩放为高100 px：`resize,h_100`
    - 缩放模式为lfit：`m_lfit`

    图片处理的URL为http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,h_100,w_300,m_lfit

    ![break](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527171.jpg)

    处理后，原图从400 px*300 px等比缩放为133 px*100 px的图片。

- 固定宽高缩放

  需求及处理参数如下：

  - 将原图缩放成宽高100 px：`resize,h_100,w_100`
  - 缩放模式fixed：`m_fixed`

  图片处理的URL为：http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,m_fixed,h_100,w_100![宽高缩放](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527175.jpg)

  处理后，原图从400 px*300 px固定缩放为100 px*100 px的图片，图片出现变形。

- 固定宽高，自动裁剪

  需求及处理参数如下：

  - 将原图缩放成宽高100 px：`resize,h_100,w_100`
  - 缩放模式fill：`m_fill`

  图片处理的URL为：http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,m_fill,h_100,w_100

  ![自动裁剪](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527179.jpg)

  处理后，原图从400 px*300 px缩放为133 px*100 px，然后居中裁剪为100px，得到大小为100 px*100 px的缩放图。

- 固定宽高，缩放填充

  需求及处理参数如下：

  - 将原图缩放成宽高100 px：`resize,h_100,w_100`

  - 缩放模式pad：`m_pad`

    **说明**

    m_pad参数不支持原图为PNG的4颜色通道的图片。

  - 以红色填充：`color_FF0000`

  图片处理的URL为：http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,m_pad,h_100,w_100,color_FF0000

  ![填充红色](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527183.jpg)

  处理后，原图从400 px*300 px缩放为100 px*75 px，然后将h居中填充到100px，得到大小为100 px*100 px的缩放图，同时填充红色。

- 按比例缩放

  需求及处理参数如下：

  将原图缩放50%：`resize,p_50`

  图片处理的URL为：http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,p_50

  ![按比例缩放](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0317789661/p527188.jpg)

  处理后，原图从400 px*300 px等比缩放为200 px*150 px，缩放为原来的50%。

# 图片水印

为保护OSS存储的图片或文件的所有权，防止资源未经授权被复制或使用，您可以为存储的资源增加水印。

## **使用场景**

- 版权保护：为保护自己的作品不被未授权使用或复制，需要在图片上加上水印来标识版权。
- 品牌推广：企业或个人为了宣传自己的品牌或标识，会在图片、视频或文档上加上带有品牌标志或名称的水印。
- 防止篡改：在某些官方文件、证书或报告上添加水印，可以增加篡改的难度，减少文件被伪造的风险。
- 抵制盗图：在网络环境中，图片很容易被他人下载和再次发布。加水印可以作为一种警示，减少他人直接盗用图片的情况。
- 法律要求：某些情况下，法律或合同条款可能要求在特定内容发布时必须加上水印，以符合规定。

## 注意事项

- 文字水印不要包号/=_#等字符
- text 文字最后使用`Urlencode`进行编码

## 参数说明

操作名称：**watermark**

相关参数如下：

- 基础参数

  | **参数**    | **是否必须** | **描述**                                                     | **取值范围**                                                 |
  | ----------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
  | **t**       | 否           | 指定图片水印或水印文字的透明度。                             | [0,100]默认值：100， 表示透明度100%（不透明）。              |
  | **g**       | 否           | 指定水印在图片中的位置。                                     | nw：左上north：中上ne：右上west：左中center：中部east：右中sw：左下south：中下se（默认值）：右下详情请参见下方基准点图片。 |
  | **x**       | 否           | 指定水印的水平边距， 即距离图片边缘的水平距离。这个参数只有当水印位置是左上、左中、左下、右上、右中、右下才有意义。 | [0,4096]默认值：10单位：像素（px）                           |
  | **y**       | 否           | 指定水印的垂直边距，即距离图片边缘的垂直距离， 这个参数只有当水印位置是左上、中上、右上、左下、中下、右下才有意义。 | [0,4096]默认值：10单位：像素（px）                           |
  | **voffset** | 否           | 指定水印的中线垂直偏移。当水印位置在左中、中部、右中时，可以指定水印位置根据中线往上或者往下偏移。 | [-1000,1000]默认值：0单位：像素（px）                        |

  水平边距、垂直边距、中线垂直偏移不仅可以调节水印在图片中的位置，当图片存在多重水印时，还可以调节水印在图中的布局。

  区域数值以及每个区域对应的基准点如下图所示。

  ![origin](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/2252359951/p2648.png)

- 文字水印参数

  | **参数**  | **是否必须** | **描述**                                              | **取值范围**                                                 |
  | --------- | ------------ | ----------------------------------------------------- | ------------------------------------------------------------ |
  | **text**  | 是           | 指定文字水印的文字内容，文字内容需进行urlencode编码。 | 最大字节长度为64个字符。                                     |
  | **color** | 否           | 指定文字水印的文字颜色，参数值为RGB颜色值。           | RGB颜色值，例如：000000表示黑色，FFFFFF表示白色。默认值：000000（黑色） |
  | **size**  | 否           | 指定文字水印的文字大小。                              | (0,1000]默认值：40单位：px                                   |

## 示例一：添加文字水印

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529184.jpg)

为example.jpg图片添加文字水印示例如下：

- 快速添加Hello World的文字水印

  对文字水印的内容Hello World进行图片处理URL为http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,w_300,h_300/watermark,text_Hello。

  ![Hello World](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/5929730761/p529185.jpg)

- 添加文字水印时配置多个图片处理参数

  为example.jpg图片添加Hello World的文字水印的同时，需要对水印文字以及原图做如下相应处理：

  - 将example.jpg缩略为宽高300：`resize,w_300,h_300`
  - 水印内容为“Hello World”：`text_Hello%20World
  - 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30`
  - 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10`

  图片处理的URL为：http://127.0.0.1:8080/file/example.png?x-oss-process=image/resize,w_300,h_300/watermark,text_Hello%20World,size_30,color_FFFFFF,t_100,g_se,x_10,y_10

  ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529186.jpg)

# 自定义裁剪

如果您希望获取指定大小的OSS图片，以适配网页元素或者满足排版格式的要求，您可以使用自定义裁剪功能。

## **使用场景**

- 网页设计与制作：在设计网页布局时，可能需要将图片裁剪成特定尺寸以适应网页元素，如头像、背景图、产品展示图等。
- 社交媒体发布：不同社交媒体平台对图片上传有各自的尺寸要求，例如封面照片、帖子图片、故事图片等，您需要按照推荐尺寸进行图片裁剪，以达到最佳展示效果。
- 移动应用开发：App中的图标、启动页、内嵌图片等都需要按规格裁剪，确保在不同分辨率和屏幕尺寸的设备上都能正确显示。
- 图像数据库管理：对于拥有大量图像资源的机构，例如图书馆、档案馆等，整理和归档时可能需要统一裁剪图片至预设尺寸。

## 注意事项

- 如果从起点开始指定的宽度和高度超过了原图，将会直接裁剪到原图边界为止。

## 参数说明

操作名称：**crop**

参数说明如下：

| **参数** | **描述**                                                     | **取值范围**                                                 |
| -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **w**    | 指定裁剪宽度。                                               | [0,图片宽度]默认为最大值。                                   |
| **h**    | 指定裁剪高度。                                               | [0,图片高度]默认为最大值。                                   |
| **x**    | 指定裁剪起点横坐标（默认左上角为原点）。                     | [0,图片边界]                                                 |
| **y**    | 指定裁剪起点纵坐标（默认左上角为原点）。                     | [0,图片边界]                                                 |
| **g**    | 设置裁剪的原点位置。原点按照九宫格的形式分布，一共有九个位置可以设置。 | nw：左上north：中上ne：右上west：左中center：中部east：右中sw：左下south：中下se：右下 (使用此参数后x,y失效) |

各裁剪原点位置的计算方法如下。其中srcW代表原图宽度，srcH代表原图高度。

| **裁剪原点** | **位置计算方法**           |
| ------------ | -------------------------- |
| nw           | 0, 0                       |
| north        | srcW/2 - w/2, 0            |
| ne           | srcW - w, 0                |
| west         | 0, srcH/2 - h/2            |
| center       | srcW/2 - w/2, srcH/2 - h/2 |
| east         | srcW - w, srcH/2 - h/2     |
| sw           | 0, srcH - h                |
| south        | srcW/2 - w/2, srcH - h     |
| se           | srcW - w, srcH - h         |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6661894861/p674595.jpg)

- 从（800,50）开始，裁减至图片边界

  需求及处理参数如下：

  - 裁剪起点为（800,50）：`crop,x_800,y_50`
  - 裁减至图片边界：裁剪时默认使用w和h的最大值，所以可省略w和h参数。

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,x_800,y_50![裁剪1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/5661894861/p674602.jpg)

- 从（800，500）开始，裁剪300 px*300 px大小的图片

  需求及处理参数如下：

  - 裁剪起点为（800,500）：`crop,x_800,y_500`
  - 裁减范围300 px*300 px：`w_300,h_300`

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,x_800,y_500,w_300,h_300

  ![裁剪2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674612.jpg)

- 裁剪原图右下角900 px*900 px的范围

  需求及处理参数如下：

  - 裁剪起点为原图右下角：`crop,g_se`
  - 裁减范围900 px*900 px：`w_900,h_900`

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,w_900,h_900,g_se

  ![裁剪3](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674614.jpg)

- 裁剪原图右下角900 px*900 px的范围，起点为相对右下九宫格的左上顶点再位移（100,100）

  需求及处理参数如下：

  - 起点为原图右下角再位移（100,100）：`crop,g_se,x_100,y_100`
  - 裁减范围900 px*900 px：`w_900,h_900`

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,x_100,y_100,w_900,h_900,g_se

  ![裁剪4](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1965894861/p674615.jpg)

# 质量变换

质量变换操作是使用原图本身的格式对图片进行压缩。您可以通过质量变换参数，修改存储在OSS内原图的质量。本文介绍对图片进行质量变换时所用到的参数及示例。

质量变换仅支持JPG和WebP，其他图片格式不支持。

## 参数说明

操作名称：**quality**

参数说明如下：

| **参数** | **描述**                                                     | **取值范围** |
| -------- | ------------------------------------------------------------ | ------------ |
| **q**    | 设置图片的相对质量，对原图按百分比进行质量压缩。例如原图质量为100%，添加`quality,q_90`参数会得到质量为90％的图片。原图质量为80%，添加`quality,q_90`参数会得到质量72%的图片。**说明**只有JPG格式的原图添加该参数，才可以决定图片的相对质量。如果原图为WebP格式，添加该参数相当于指定了原图绝对质量，即与参数**Q**的作用相同。 | [1,100]      |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7442799661/p529275.jpg)



- 变换图片相对质量

  需求及处理参数如下：

  - 原图缩放为宽100 px：`resize,w_100`
  - 图片相对质量设置为80%：`quality,q_80`

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_100/quality,q_80

  ![变换1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8442799661/p529279.jpg)

# 格式转换

通过格式转换参数，您无需将图片下载到本地进行转换，只需指定URL转换存储在OSS内原图的格式。本文介绍对图片进行格式转换时所用到的参数及示例。

## **使用场景**

- 适应不同设备和平台：不同的浏览器、操作系统或移动设备可能支持不同的图片格式。例如，WebP格式在浏览器上能提供更好的压缩效率，通过OSS图片格式转换，可以将上传的图片转换为多种格式，确保在各种终端上的兼容性和最佳显示效果。
- 节省存储成本：某些图片格式（如WebP）在保证视觉质量的前提下，文件大小通常小于其他传统格式。通过格式转换功能，可以在不牺牲图像质量的情况下减少存储空间占用，从而降低存储成本。
- 统一资源管理：在电商、社交、媒体等行业中，需要对大量用户上传的图片进行标准化处理。您可以将上传的图片统一转换成指定格式，便于后续的一致管理和分发。

## 注意事项

- 图片处理包含缩放操作时，建议将格式转换参数放到处理参数的最后。

  例如`image/resize,w_100/format,jpg`

- 图片处理包含缩放和水印操作时，建议将格式转换参数添加在缩放参数之后。

  例如`image/resize,w_100/format,jpg/watermark,...`

- 如果原图没有透明通道，转换成PNG、Web、BMP等存在透明通道的格式，默认会把透明填充成白色。

- OSS不支持将透明色填充为黑色。

## 参数说明

操作名称：**format**

参数说明如下：

| **取值范围** | **描述**                                                     |
| ------------ | ------------------------------------------------------------ |
| **jpg**      | 将原图保存为JPG格式。**重要**不支持将存在透明通道的HEIC格式的图片保存为JPG格式。 |
| **png**      | 将原图保存为PNG格式。                                        |
| **webp**     | 将原图保存为WebP格式。                                       |
| **bmp**      | 将原图保存为BMP格式。                                        |
| **gif**      | 原图为GIF图片则继续保存为GIF格式；原图不是GIF图片，则按原图格式保存。 |
| **tiff**     | 将原图保存为TIFF格式。                                       |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.gif，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.gif

![gif](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139212.png)



- 将原图转换为PNG格式

  图片处理URL为：http://127.0.0.1:8080/file/example.gif?x-oss-process=image/format,png![png](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8448459951/p139213.png)

- 将原图转换成JPG格式，并支持渐进显示

  需求及处理参数如下：

  - 图片设置为渐进显示：`interlace,1`
  - 图片转换为JPG格式：`format,jpg`

  图片处理URL为：http://127.0.0.1:8080/file/example.gif?x-oss-process=image/format,jpg

  ![img](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0548459951/p2555.jpg)

- 将原图缩放为宽200 px，并转换为WebP格式

  需求及处理参数如下：

  - 图片缩放为宽200 px：`resize,w_200`
  - 图片转换为WebP格式：`format,webp`

  图片处理URL为：http://127.0.0.1:8080/file/example.gif?x-oss-process=image/resize,w_200,h_100/format,webp

目前转换为webp还不支持动图。

# 模糊效果

如果您希望保护OSS图片的隐私信息或者提升图片的视觉效果，您可以对OSS图片增加全局或者局部模糊效果。

## **使用场景**

- 保护隐私信息：在发布包含敏感信息的图片前保护隐私信息。
- 多图层合成：在多图层合成过程中，适当使用模糊效果可以平滑不同图层间的边缘，提供更舒适的视觉体验。
- 低分辨率掩饰：当图片原始分辨率较低，无法满足高清展示需求时，适度的模糊处理能够减轻像素感。

## 参数说明

操作名称：**blur**

Blur 使用高斯函数生成图像的模糊版本。
参数必须为正数，表示图像的模糊程度。

参数说明如下：

| **参数**    | **是否必须** | **描述**         | **取值范围**                 |
| ----------- | ------------ | ---------------- | ---------------------------- |
| **[value]** | 是           | 设置高斯函数参数 | [1,50]该值越大，图片越模糊。 |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6661894861/p674595.jpg)

- 模糊图片

  需求及处理参数为：高斯参数为5的模糊处理`5`。

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/blur,5![模糊1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/9692994861/p674663.jpg)

# 旋转

您可以通过旋转参数，将存储在OSS内的原图按指定方向旋转。本文介绍旋转图片时所用到的参数和示例。

## 参数说明

操作名称：**rotate**

参数说明如下：

| **参数**    | **描述**                 | **取值范围**                   |
| ----------- | ------------------------ | ------------------------------ |
| **[value]** | 图片按顺时针旋转的角度。 | [0,360]默认值：0，表示不旋转。 |

## 注意事项

- 若图片旋转的角度不是90°、180°、270°、360°时，会导致处理后的图片尺寸变大。
- 旋转功能对图片的尺寸有限制，图片的宽或者高不能超过4096 px。
- GIF图片不支持旋转参数。如果对GIF动图执行旋转操作，GIF动图会变成GIF静态图。

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529611.jpg)

- 将原图按顺时针旋转90°

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/rotate,90

  ![旋转1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/0212120761/p529612.jpg)

- 将原图按顺时针旋转70°

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/rotate,70

  # 亮度

  您可以通过亮度参数，调节存储在OSS内的原图亮度。本文介绍调节图片亮度时所用到的参数及示例。

  ## 参数说明

  操作名称：**bright**

  参数说明如下：

  | **参数**    | **描述**         | **取值范围**                                                 |
  | ----------- | ---------------- | ------------------------------------------------------------ |
  | **[value]** | 指定图片的亮度。 | [-100, 100]取值＜0：降低图片亮度。取值=0：不调整图片亮度。取值＞0：提高图片亮度。 |

  ## 示例

  您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，图片访问地址为：

  http://127.0.0.1:8080/file/example.jpg![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8532220761/p529913.jpg)

  - 将图片亮度提高50

    图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/bright,50![亮度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7532220761/p529917.jpg)

  - 将图片亮度降低50

    图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/bright,-50

    

# 锐化

您可以通过锐化参数，提高存储在OSS内原图的清晰度。本文介绍对图片进行锐化时所用到的参数及示例。

## 参数说明

操作名称：**sharpen**

参数说明如下：

| **参数**    | **描述**             | **取值范围**                                                 |
| ----------- | -------------------- | ------------------------------------------------------------ |
| **[value]** | 设置锐化效果的强度。 | [50,399]取值越大，图片越清晰，但过大的值可能会导致图片失真。为达到较优效果，推荐取值为100。 |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529928.jpg)



对原图进行锐化处理，锐化参数为100。图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/sharpen,100![锐化1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/1162220761/p529929.jpg)

# 对比度

对比度是指一幅图像中明暗区域最亮的白和最暗的黑之间不同亮度层级的测量，即指一幅图像灰度反差的大小。您可以通过对比度参数，调整存储在OSS内原图的对比度。本文介绍调节图片对比度时所用到的参数及示例。

## 参数说明

操作名称：**contrast**

| **参数**    | **描述**           | **取值范围**                                                 |
| ----------- | ------------------ | ------------------------------------------------------------ |
| **[value]** | 指定图片的对比度。 | [-100,100]取值＜0：降低图片对比度。取值=0：维持原图对比度。取值＞0：提高图片对比度。 |

## 示例

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例]: #3.综合示例

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529934.jpg)

- 对比度降低50

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/contrast,-50

  ![对比度1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529935.jpg)

- 对比度提高50

  图片处理URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/contrast,50

![对比度2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/8782220761/p529938.jpg)