package imageprocess

import (
	"image"
	"image/gif"
	"io"
	"strings"
)

func Process(img image.Image, w io.Writer, originalFormat Format, options []Option) error {
	//循环 options 处理
	img, format, quality, err := processImage(img, options)
	if err != nil {
		return err
	}
	//format and quality
	if format == "" {
		format = originalFormat
	}
	if quality == 0 {
		quality = 100
	}
	return EncodeImage(img, w, format, quality)
}

func ProcessGif(gifImg *gif.GIF, w io.Writer, options []Option) error {
	isGif := true
	parameterTypes := getParameterTypes(options)
	if contains(parameterTypes, string(FormatType)) {
		parmsStr := SerializeOptions(options)
		if !strings.Contains(parmsStr, "gif") {
			// 不转换为gif
			isGif = false
		}
	}
	// 是否转换为GIF
	if isGif {
		// 遍历GIF的每一帧进行缩放
		for i := range gifImg.Image {
			original := gifImg.Image[i]
			// 使用nfnt/resize库对图像进行缩放
			resized, _, _, err := processImage(original, options)
			if err != nil {
				return err
			}
			// 将缩放后的图像替换掉原图像
			gifImg.Image[i] = image.NewPaletted(resized.Bounds(), original.Palette)
			for y := 0; y < resized.Bounds().Dy(); y++ {
				for x := 0; x < resized.Bounds().Dx(); x++ {
					gifImg.Image[i].Set(x, y, resized.At(x, y))
				}
			}
			// 更新每一帧的尺寸
			gifImg.Config.Width = int(resized.Bounds().Dx())
			gifImg.Config.Height = int(resized.Bounds().Dy())
		}
		return gif.EncodeAll(w, gifImg)
	} else {
		return Process(gifImg.Image[0], w, GIF, options)
	}
}

func processImage(img image.Image, options []Option) (image.Image, Format, int, error) {
	var format Format
	var quality int
	for _, option := range options {
		switch option.Parameter {
		case Resize:
			img = ResizeImage(img, option.Option.(ResizeOption))
		case Crop:
			img = CropImage(img, option.Option.(CropOption))
		case Watermark:
			img = WarterMarkText(img, option.Option.(TextWatermarkOption))
		case Rotate:
			img = AddjustRotate(img, option.Option.(RotateOption))
		case Blur:
			img = AdjustBlur(img, option.Option.(BlurOption))
		case Sharpen:
			img = AddjustSharpen(img, option.Option.(SharpenOption))
		case Saturation:
			img = AddjustSaturation(img, option.Option.(SaturationOption))
		case Gamma:
			img = AddjustGamma(img, option.Option.(GammaOption))
		case Brightness:
			img = AddjustBright(img, option.Option.(BrightnessOption))
		case Contrast:
			img = AddjustContrast(img, option.Option.(ContrastOption))
		case FormatType:
			format = option.Option.(FormatOption).Format
		case Quality:
			quality = option.Option.(QualityOption).Quality
		default:
			return nil, "", 100, ErrUnsupportedParameter
		}
	}
	return img, format, quality, nil
}

// getParameterTypes 接收一个Option类型的切片，并返回一个Parameter类型的切片。
// 该函数的目的是从提供的Option中提取出所有的parameter。
//
// 参数:
//
//	options - 一个包含多个Option的切片，每个Option都嵌有一个Parameter。
//
// 返回值:
//
//	一个Parameter类型的切片，包含了从options中提取出的所有parameter。
func getParameterTypes(options []Option) []string {
	parameterTypes := make([]string, 0)
	for _, option := range options {
		parameterTypes = append(parameterTypes, string(option.Parameter))
	}
	return parameterTypes
}

func contains(arr []string, target string) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}
	return false
}
