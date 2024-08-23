package imageprocess

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"net/http"
	"os"
	"strings"
)

// parseHexColor converts a hex color string to an image/color.RGBA.
func ParseHexColor(s string) (*color.RGBA, error) {
	c := color.RGBA{A: 0xff} // Default to opaque
	if len(s) != 6 {
		return &c, nil
	}
	_, err := fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	if err != nil {
		return nil, fmt.Errorf("invalid color format: %s, err: %s", s, err.Error())
	}
	return &c, nil
}

func ParseRGBAColor(c *color.RGBA) string {
	return fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
}

// downloadImage downloads an image from a URL and returns it as an image.Image.
func DownloadImage(url string) (image.Image, string, error) {
	// Send the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("failed to fetch image: status code %d", resp.StatusCode)
	}

	// Decode the image
	img, fileName, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fileName, err
	}

	return img, fileName, nil
}

// LoadImage loads an image from the specified file path.
func LoadImage(filePath string) (image.Image, Format, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()
	fileExt := strings.ToLower(filePath[strings.LastIndex(filePath, ".")+1:])

	// 检查文件扩展名是否为支持的图像格式
	format, err := FormatFromExtension(fileExt)
	if err != nil {
		return nil, format, fmt.Errorf("unsupported image format: %s", fileExt)
	}
	img, err := DecodeImage(file, format)
	return img, format, err
}
func LoadGif(filePath string) (*gif.GIF, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %v", err)
	}
	defer file.Close()
	img, err := gif.DecodeAll(file)
	return img, err
}

/*
获取位置起始点
option.X 指定水印的水平边距， 即距离图片边缘的水平距离
option.Y 指定水印的垂直边距， 即距离图片边缘的垂直距离
w 水印的宽度
h 水印的高度
。_______________X(W)
|***(0,h)
|
|
|
|****(0,Ymax)
Y(H)
*/
func TextWatermarkPositionPoint(img image.Image, option TextWatermarkOption, w, h float64) (x, y float64) {
	rg := img.Bounds()
	srcW := float64(rg.Dx())
	srcH := float64(rg.Dy())
	switch option.Position {
	case NorthWest:
		return float64(option.X), h + float64(option.Y)
	case North:
		return srcW/2 - w/2, h + float64(option.Y)
	case NorthEast:
		return srcW - w - float64(option.X), h + float64(option.Y)
	case West:
		return float64(option.X), srcH/2 + h
	case Center:
		return srcW/2 - w/2, srcH/2 + h
	case East:
		return srcW - w - float64(option.X), srcH/2 + h
	case SouthWest:
		return float64(option.X), srcH - float64(option.Y)
	case South:
		return srcW/2 - w/2, srcH - float64(option.Y)
	case SouthEast:
		return srcW - w - float64(option.X), srcH - float64(option.Y)
	default:
		return 0, 0
	}
}

/*
裁剪原点
位置计算方法
*/
func CropPositionPoint(img image.Image, node Position, w, h float64) (x, y float64) {
	rg := img.Bounds()
	srcW := float64(rg.Dx())
	srcH := float64(rg.Dy())
	switch node {
	case NorthWest:
		return 0, 0
	case North:
		return srcW/2 - w/2, 0
	case NorthEast:
		return srcW - w, 0
	case West:
		return 0, srcH/2 - h/2
	case Center:
		return srcW/2 - w/2, srcH/2 - h/2
	case East:
		return srcW - w, srcH/2 - h/2
	case SouthWest:
		return 0, srcH - h
	case South:
		return srcW/2 - w/2, srcH - h
	case SouthEast:
		return srcW - w, srcH - h
	default:
		return 0, 0
	}
}
