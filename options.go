package imageprocess

import (
	"fmt"
	"image/color"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

type ResizeMode string

const (
	Lfit  ResizeMode = "lfit" // default
	Mfit  ResizeMode = "mfit"
	Fill  ResizeMode = "fill"
	Pad   ResizeMode = "pad"
	Fixed ResizeMode = "fixed"
)

// Supported formats
var supportedResizeModes = map[string]ResizeMode{
	"lfit":  Lfit,
	"mfit":  Mfit,
	"fill":  Fill,
	"pad":   Pad,
	"fixed": Fixed,
}

/*
*
*
| **Parameter** | **Required**                          | **Description**                                              | **Value range**                                              |
| ------------- | ------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **m**         | Yes                                   | Specifies the type of the resize action.                     | lfit: OSS proportionally resizes the source image as large as possible in a rectangle based on the specified width and height. This is the default value.mfit: OSS proportionally resizes the source image as small as possible outside a rectangle based on the specified width and height.fill: OSS proportionally resizes the source image as small as possible outside a rectangle, and then crops the resized image from the center based on the specified width and height.pad: OSS resizes the source image as large as possible in a rectangle based on the specified width and height, and fills the empty space with a specific color.fixed: OSS forcibly resizes the source image based on the specified width and height.For more information about examples of each resizing type, see [Examples](https://www.alibabacloud.com/help/en/oss/user-guide/resize-images-4?spm=a2c63.p38356.0.0.1ba3f386WmEBnU#li-llf-951-r0b).**Important**If you set this parameter to lfit or mfit, the aspect ratio (the ratio of the width to the height) of the source image is rounded to an integer if the ratio is a decimal.If you specify the m parameter and specify a value for w or h, the values specified for l and s do not take effect. |
| **w**         | No                                    | Specifies the width to which you want to resize the image.   | [1,16384]                                                    |
| **h**         | No                                    | Specifies the height to which you want to resize the image.  | [1,16384]                                                    |
| **l**         | Yes                                   | Specifies the length of the longer side to which you want to resize the image.**Note**The longer side is the side for which the ratio of the source length to the target length is larger. For example, if a source image is resized from 400 × 200 pixels to 800 × 100 pixels, the source-to-target size ratios are 0.5 (400/800) and 2 (200/100). 0.5 is smaller than 2. As a result, the side that contains 200 pixels is used as the longer side. | [1,16384]                                                    |
| **s**         | Yes                                   | Specifies the length of the shorter side to which you want to resize the image.**Note**The shorter side is the side for which the ratio of the source length to the target length is smaller. For example, if a source image is resized from 400 × 200 pixels to 800 × 100 pixels, the source-to-target ratios are 0.5 (400/800) and 2 (200/100). 0.5 is smaller than 2. As a result, the side that contains 400 pixels is used as the shorter side. | [1,16384]                                                    |
| **limit**     | No                                    | Specifies whether to resize the image when the resolution of the target image is higher than the resolution of the source image. | 1: This is the default value. OSS returns the image that is processed based on the resolution of the source image. The size of the returned image may be different from the size of the source image.0: OSS resizes the source image based on the specified value.**Note**The size of a GIF image can only be reduced. |
| **color**     | Yes (only when the value of m is pad) | If you set the resizing type to pad, you can select a color to fill the empty space. | RGB color values. For example, 000000 indicates black, and FFFFFF indicates white.Default value: FFFFFF (white). |
*/
type ResizeOption struct {
	ResizeMode ResizeMode  `option:"m"`
	Width      int         `option:"w,optional"`
	Height     int         `option:"h,optional"`
	Longer     int         `option:"l"` // 不适用
	Shorter    int         `option:"s"` // 不适用
	Limit      int         `option:"limit"`
	Color      *color.RGBA `option:"color"`
}

type Position string

const (
	NorthWest Position = "nw"
	North     Position = "north"
	NorthEast Position = "ne"
	West      Position = "west"
	Center    Position = "center"
	East      Position = "east"
	SouthWest Position = "sw"
	South     Position = "south"
	SouthEast Position = "se" // default
)

var supportedPositions = map[string]Position{
	"nw":     NorthWest,
	"north":  North,
	"ne":     NorthEast,
	"west":   West,
	"center": Center,
	"east":   East,
	"sw":     SouthWest,
	"south":  South,
	"se":     SouthEast,
}

/*
*
| **Parameter** | **Required** | **Description**                                              | **Valid value**                                              |
| ------------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **t**         | No           | The opacity of the watermark.                                | [0,100]Default value: 100. The value 100 specifies that the watermark is opaque. |
| **g**         | No           | The position of the watermark on the image.                  | nw: upper left.north: upper middle.ne: upper right.west: middle left.center: center.east: middle right.sw: lower left.south: lower middle.se: lower right. This is the default value.For the precise position that is specified by each value, see the following figure. |
| **x**         | No           | The horizontal margin, which indicates the horizontal distance between the watermark and the image edge. This parameter takes effect only when the watermark is on the upper left, middle left, lower left, upper right, middle right, or lower right of the image. | [0,4096]Default value: 10.Unit: pixel.                       |
| **y**         | No           | The vertical margin that specifies the vertical distance between the watermark and the image edge. This parameter takes effect only when the watermark is on the upper left, upper middle, upper right, lower left, lower middle, or lower right of the image. | [0,4096]Default value: 10.Unit: pixel.                       |
| **voffset**   | No           | The vertical offset from the middle line. When the watermark is on the middle left, center, or middle right of the image, you can specify the vertical offset of the watermark along the middle line. | [-1000,1000]Default value: 0.Unit: pixel.                    |
| **fill**      | No           | Specifies whether to tile the image watermarks or text watermarks across the image.**Note**If you want to add tiled watermarks, submit an application at [Quota Center](https://oss.console.aliyun.com/more-tool/quota-setting).. | 1: tiles the image watermarks or text watermarks across the image.0: does not tile the image watermarks or text watermarks across the image. This is the default value. |
| **padx**      | No           | The horizontal spacing between watermarks when the image watermarks or text watermarks are tiled across the image. This parameter is valid only when you set fill to 1. | [0,4096]Default value: 0.Unit: pixel.                        |
| **pady**      | No           | The vertical spacing between watermarks when the image watermarks or text watermarks are tiled across the image. This parameter is valid only when you set fill to 1. | [0,4096]Default value: 0.Unit: pixel.                        |
*/
type WatermarkOption struct {
	Opacity  int      `option:"t,optional"`
	Position Position `option:"g,optional"`
	X        int      `option:"x,optional"`
	Y        int      `option:"y,optional"`
	Voffset  int      `option:"voffset,optional"`
	Fill     int      `option:"fill,optional"`
	Padx     int      `option:"padx,optional"`
	Pady     int      `option:"pady,optional"`
}

/*
*
| **Parameter** | **Required** | **Description**                                              | **Valid value**                                              |
| ------------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **text**      | Yes          | The content of the text watermark. The text content must be Base64-encoded. For more information, see [Encode watermark-related parameters](https://www.alibabacloud.com/help/en/oss/user-guide/add-watermarks?spm=a2c63.p38356.0.0.297b1a29rh4w2M#watermark). | A Chinese string before Base64-encoding cannot exceed 64 characters in length. |
| **type**      | No           | The font of the text watermark. The font name must be Base64-encoded. | For more information about the supported fonts and the encoding results of the fonts, see [Font types and encoding results](https://www.alibabacloud.com/help/en/oss/user-guide/add-watermarks?spm=a2c63.p38356.0.0.297b1a29rh4w2M#table-ipy-1vu-isp).Default value: wqy-zenhei (encoding result: d3F5LXplbmhlaQ). |
| **color**     | No           | The color of the text watermark. The valid values for this parameter are RGB color values. | For example, 000000 specifies black, and FFFFFF specifies white.Default value: 000000. |
| **size**      | No           | The size of the text watermark.                              | (0,1000]Default value: 40.Unit: pixel.                       |
| **shadow**    | No           | The opacity of the shadow for the text watermark.            | [0,100]Default value: 0. The value 0 specifies that no shadows are added to the text. |
| **rotate**    | No           | The degree by which the text is rotated clockwise.           | [0,360]Default value: 0. The value 0 specifies that the text is not rotated. |
*/

type TextWatermarkOption struct {
	WatermarkOption
	Text   string      `option:"text"`
	Type   string      `option:"type,optional"`
	Color  *color.RGBA `option:"color,optional"`
	Size   int         `option:"size,optional"`
	Shadow int         `option:"shadow,optional"`
	Rotate int         `option:"rotate,optional"`
}

// not support
const (
	Face Position = "face"
	Auto Position = "auto"
)

// not support
// var cropPositionMap = map[string]Position{
// 	"face": Face,
// 	"auto": Auto,
// }

/*
*
| **Parameter** | **Description**                                              | **Value range**                                              |
| ------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **w**         | The width that you want to crop.                             | [0, image width]Default value: the maximum value.            |
| **h**         | The height that you want to crop.                            | [0, image height]Default value: the maximum value.           |
| **x**         | The X coordinate of the area that you want to crop. The default value is the X coordinate of the upper-left corner of the image. | [0, image bound]                                             |
| **y**         | The Y coordinate of the area that you want to crop. The default value is the Y coordinate of the upper-left corner of the image. | [0, image bound]                                             |
| **g**         | The position of the area that you want to crop in a 3 x 3 grid. The image is located in a 3 x 3 grid. The grid has nine tiles. | nw: upper leftnorth: upper middlene: upper rightwest: middle leftcenter: centereast: middle rightsw: lower leftsouth: lower middlese: lower rightFor more information about how to calculate the position of each tile, see [the following table](https://www.alibabacloud.com/help/en/oss/user-guide/custom-crop?spm=a2c63.p38356.0.0.7d7d6661kGXVLR#table-xdo-tzc-rfu). |
*/
type CropOption struct {
	Width    int      `option:"w,optional"`
	Height   int      `option:"h,optional"`
	X        int      `option:"x,optional"`
	Y        int      `option:"y,optional"`
	Position Position `option:"g,optional"`
}

/*
*
Imaging supports image resizing using various resampling filters. The most notable ones:

Lanczos - A high-quality resampling filter for photographic images yielding sharp results.
CatmullRom - A sharp cubic filter that is faster than Lanczos filter while providing similar results.
MitchellNetravali - A cubic filter that produces smoother results with less ringing artifacts than CatmullRom.
Linear - Bilinear resampling filter, produces smooth output. Faster than cubic filters.
Box - Simple and fast averaging filter appropriate for downscaling. When upscaling it's similar to NearestNeighbor.
NearestNeighbor - Fastest resampling filter, no antialiasing.
The full list of supported filters: NearestNeighbor, Box, Linear, Hermite, MitchellNetravali, CatmullRom, BSpline, Gaussian, Lanczos, Hann, Hamming, Blackman, Bartlett, Welch, Cosine. Custom filters can be created using ResampleFilter struct.
*/
const (
	Lanczos         = "lanczos"
	CatmullRom      = "catmullrom"
	Linear          = "linear"
	Box             = "box"
	NearestNeighbor = "nearestneighbor"
)

var resampleFilterMap = map[string]imaging.ResampleFilter{
	Lanczos:         imaging.Lanczos,
	CatmullRom:      imaging.CatmullRom,
	Linear:          imaging.Linear,
	Box:             imaging.Box,
	NearestNeighbor: imaging.NearestNeighbor,
}

/*
| **Parameter** | **Description**                                              | **Value range**                                              |
| ------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **q**         | The quality of the image.                                    | [0,100]Default value: 100.                                   |
*/
type QualityOption struct {
	Quality int                    `option:"q,optional"`
	Fiter   imaging.ResampleFilter `option:"f,optional"`
}

/*
*
| **Valid value** | **Description**                                              |
| --------------- | ------------------------------------------------------------ |
| **jpg**         | Converts the format of a source image to JPG.**Important**Images in the HEIC format that support alpha channels cannot be converted to JPG images. |
| **png**         | Converts the format of a source image to PNG.                |
| **webp**        | Converts the format of a source image to WebP.               |
| **bmp**         | Converts the format of a source image to BMP.                |
| **gif**         | Converts the format of a source image to GIF. The conversion takes effect only when the source image is also a GIF image. If the source image is not in the GIF format, the processed image is stored in the original format. |
| **tiff**        | Converts the format of a source image to TIFF.               |
*/
type FormatOption struct {
	Format Format `option:",optional"`
}

type BlurOption struct {
	Radius int `option:"r,optional"`
}
type BrightnessOption struct {
	Value int `option:",optional"`
}
type SharpenOption struct {
	Value int `option:",optional"`
}
type ContrastOption struct {
	Value int `option:",optional"`
}
type GammaOption struct {
	Value float64 `option:",optional"`
}
type SaturationOption struct {
	Value int `option:",optional"`
}
type HueOption struct {
	Value int `option:",optional"`
}

/*
*
The degree by which the image is rotated clockwise.
*/
type RotateOption struct {
	Value int `option:",optional"`
}

type Parameter string

const (
	Resize     Parameter = "resize"
	Watermark  Parameter = "watermark"
	FormatType Parameter = "format"
	Crop       Parameter = "crop"
	Blur       Parameter = "blur"
	Brightness Parameter = "bright"
	Sharpen    Parameter = "sharpen"
	Contrast   Parameter = "contrast"
	Gamma      Parameter = "gamma"
	Saturation Parameter = "saturation"
	Hue        Parameter = "hue"
	Rotate     Parameter = "rotate"
	Quality    Parameter = "quality"
)

type Option struct {
	Parameter Parameter   `option:"parameter"`
	Option    interface{} `option:"option"`
}

// parseUrl解析给定的URL字符串，提取并解析其中的"x-oss-process"参数。
// 参数urlstr是需要解析的URL字符串。
// 返回值是一个映射，其中包含解析后的参数，以及一个错误对象。
// 如果解析过程中发生错误，会返回nil和错误对象。
// example:
// https://oss-console-img-demo-cn-hangzhou-3az.oss-cn-hangzhou.aliyuncs.com/example.gif?x-oss-process=image/format,png
func ParseUrlOptions(urlstr string) ([]Option, error) {
	// 使用url.Parse解析URL字符串，如果解析失败，直接返回错误。
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	// 获取查询参数"x-oss-process"的值。
	parameterStr := u.Query().Get("x-oss-process")
	// 解析"x-oss-process"参数的值，并返回解析结果。
	return ParseOptions(parameterStr)
}

// 根据传入的参数构建不同的图像处理选项
// parameter: 图像处理操作的参数，格式为"操作名_参数1_参数2..."
// 返回一个映射，键为图像处理操作类型，值为相应的处理选项
// example:
// image/resize,h_100,m_lfit/format,jpg
func ParseOptions(parameterStr string) (options []Option, err error) {
	options = make([]Option, 0)
	if parameterStr == "" {
		return nil, fmt.Errorf("parameterStr is empty")
	}
	//去除前缀
	parameterStr = strings.TrimPrefix(parameterStr, "image/")
	//将 parameterStr 解析成Option
	flag := false
	for _, parameter := range strings.Split(parameterStr, "/") {
		params := strings.Split(parameter, ",")
		switch Parameter(params[0]) {
		case Resize:
			if len(params) > 1 {
				rp := ResizeOption{ResizeMode: Lfit, Limit: 1, Color: &color.RGBA{R: 255, G: 255, B: 255, A: 255}}
				for i := 1; i < len(params); i++ {
					kv := strings.Split(params[i], "_")
					switch Parameter(kv[0]) {
					case "w":
						rp.Width, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "h":
						rp.Height, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "m":
						if f, ok := supportedResizeModes[kv[1]]; ok {
							rp.ResizeMode = f
							flag = true
						} else {
							return nil, fmt.Errorf("resize mode('m') is not supported")
						}
					case "l":
						rp.Longer, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "s":
						rp.Shorter, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "color":
						rp.Color, err = ParseHexColor(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a color:%v", kv[0], err)
						}
						flag = true
					case "limit":
						rp.Limit, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					default:
						return nil, fmt.Errorf("resize parameter('%v') is not supported", kv[0])
					}
				}
				if rp.Width == 0 || rp.Height == 0 {
					return nil, fmt.Errorf("w and h is cannot be empty")
				}
				if flag {
					options = append(options, Option{Resize, rp})
				}
			}
		case Watermark:
			if len(params) > 1 {
				wp := TextWatermarkOption{WatermarkOption: WatermarkOption{
					Opacity: 100, Position: SouthEast, X: 10, Y: 10,
				}, Color: &color.RGBA{0, 0, 0, 1}, Size: 40}
				for i := 1; i < len(params); i++ {
					kv := strings.Split(params[i], "_")
					switch Parameter(kv[0]) {
					case "x":
						wp.X, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "y":
						wp.Y, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "voffset":
						wp.Voffset, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "g":
						if f, ok := supportedPositions[kv[1]]; ok {
							wp.Position = f
							flag = true
						} else {
							return nil, fmt.Errorf("position('g') is not supported")
						}
					case "t":
						wp.Opacity, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "fill":
						wp.Fill, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "pady":
						wp.Pady, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "padx":
						wp.Padx, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "text":
						wp.Text = kv[1]
						flag = true
					case "type":
						wp.Type = kv[1]
						flag = true
					case "shadow":
						wp.Shadow, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "color":
						wp.Color, err = ParseHexColor(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a color:%v", kv[0], err)
						}
						flag = true
					case "rotate":
						wp.Rotate, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "size":
						wp.Size, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
					}
				}
				if wp.Text == "" {
					return nil, fmt.Errorf("text is cannot be empty")
				}
				if flag {
					options = append(options, Option{Watermark, wp})
				}
			}
		case Crop:
			if len(params) > 1 {
				cp := CropOption{}
				for i := 1; i < len(params); i++ {
					kv := strings.Split(params[i], "_")
					switch Parameter(kv[0]) {
					case "w":
						cp.Width, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "h":
						cp.Height, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "x":
						cp.X, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "y":
						cp.Y, err = strconv.Atoi(kv[1])
						if err != nil {
							return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
						}
						flag = true
					case "g":
						if f, ok := supportedPositions[kv[1]]; ok {
							cp.Position = f
							flag = true
						} else {
							return nil, fmt.Errorf("crop mode('m') is not supported")
						}
					}
				}
				if flag {
					options = append(options, Option{Crop, cp})
				}
			}
		case Blur:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Blur, BlurOption{Radius: v}})
			}
		case Brightness:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Brightness, BrightnessOption{Value: v}})
			}
		case Sharpen:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Sharpen, SharpenOption{Value: v}})
			}
		case Contrast:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Contrast, ContrastOption{Value: v}})
			}
		case Gamma:
			if len(params) > 1 {
				// string to float
				v, err := strconv.ParseFloat(params[1], 32)
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Gamma, GammaOption{Value: v}})
			}
		case Saturation:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Saturation, SaturationOption{Value: v}})
			}
		case Hue:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Hue, HueOption{Value: v}})
			}
		case Rotate:
			if len(params) > 1 {
				v, err := strconv.Atoi(params[1])
				if err != nil {
					return nil, fmt.Errorf("%v is not a number:%v", params[0], err)
				}
				options = append(options, Option{Rotate, RotateOption{Value: v}})
			}
		case FormatType:
			if len(params) > 1 {
				if f, ok := supportedFormats[params[1]]; ok {
					options = append(options, Option{FormatType, FormatOption{Format: f}})
					flag = true
				} else {
					return nil, fmt.Errorf("format('%v') is not supported", params[1])
				}
			}
		case Quality:
			if len(params) > 1 {
				qp := QualityOption{}
				kv := strings.Split(params[1], "_")
				switch Parameter(kv[0]) {
				case "q":
					qp.Quality, err = strconv.Atoi(kv[1])
					if err != nil {
						return nil, fmt.Errorf("%v is not a number:%v", kv[0], err)
					}
					flag = true
				case "f":
					if f, ok := resampleFilterMap[kv[1]]; ok {
						qp.Fiter = f
						flag = true
					} else {
						return nil, fmt.Errorf("%v is not supported", kv[0])
					}
				}
				if flag {
					options = append(options, Option{Quality, qp})
				}
			}
		default:
			return nil, fmt.Errorf("parameter %s is not support", parameter)
		}
	}
	return options, nil
}

// 序列化Options
func SerializeOptions(options []Option) string {
	//遍历 optionMap
	if len(options) == 0 {
		return ""
	}
	// 获取参数
	var parts []string
	for _, v := range options {
		//组装
		param := serializeOption(v.Option)
		if param == "" {
			continue
		}
		parts = append(parts, fmt.Sprintf("%s,%s", v.Parameter, param))
		//v是struct 解析字段 tag `option:"r,optional"`
	}
	return strings.Join(parts, "/")
}

func serializeOption(option interface{}) string {
	var parts []string

	// Use reflection to iterate over struct fields
	val := reflect.ValueOf(option)
	typ := reflect.TypeOf(option)

	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("option")
		tagParts := strings.Split(tag, ",")

		if len(tagParts) > 0 && field.IsValid() {
			// Convert field value to string based on its type
			var fieldStr string
			switch field.Kind() {
			case reflect.Int:
				if field.Int() == 0 {
					continue
				}
				fieldStr = strconv.Itoa(int(field.Int()))
			case reflect.Pointer:
				if field.IsNil() {
					continue
				}
				//获取属性名称
				typ := reflect.TypeOf(field.Interface()).String() == "*color.RGBA"
				if typ {
					fieldStr = ParseRGBAColor(field.Interface().(*color.RGBA))
				} else {
					fieldStr = fmt.Sprintf("%v", field.Interface())
				}
			case reflect.String:
				if field.String() == "" {
					continue
				}
				fieldStr = field.String()
			case reflect.Bool:
				if !field.Bool() {
					continue
				}
				fieldStr = strconv.FormatBool(field.Bool())
			default:
				if field.Interface() == nil {
					continue
				}
				// imageprocess.WatermarkOption jump in
				typ := reflect.TypeOf(field.Interface()).String() == "imageprocess.WatermarkOption"
				if typ {
					fieldStr = serializeOption(field.Interface())
				} else {
					fieldStr = fmt.Sprintf("%v", field.Interface())
				}
			}

			// Combine tag and field value
			if tagParts[0] == "" {
				parts = append(parts, fieldStr)
			} else {
				parts = append(parts, fmt.Sprintf("%s_%s", tagParts[0], fieldStr))
			}
		}
	}

	// Join all parts with commas
	return strings.Join(parts, ",")
}
