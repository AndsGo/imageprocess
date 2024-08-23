package imageprocess

import (
	"fmt"
	"image"
	"image/gif"

	"github.com/disintegration/imaging"
)

// ResizeImage resizes the given image according to specified parameters.
func ResizeImage(img image.Image, option ResizeOption) image.Image {
	if option.Height == 0 || option.Width == 0 {
		return img
	}
	switch option.ResizeMode {
	case Lfit:
		return imaging.Fit(img, option.Width, option.Height, imaging.Lanczos)
	case Mfit:
		return imaging.Fit(img, option.Width, option.Height, imaging.Lanczos)
	case Fill:
		return imaging.Fill(img, option.Width, option.Height, imaging.Center, imaging.Lanczos)
	case Pad:
		//将原图缩放为指定w与h的矩形内的最大图片，然后使用指定颜色居中填充空白部分。
		bg := imaging.New(option.Width, option.Height, option.Color)
		return imaging.PasteCenter(bg, imaging.Fit(img, option.Width, option.Height, imaging.Lanczos))
		// return imaging.Resize(img, width, height, imaging.Lanczos)
	case Fixed:
		return imaging.Resize(img, option.Width, option.Height, imaging.Lanczos)
	default:
		fmt.Println("Unsupported resize mode")
		return img
	}
}

func ResizeGif(gifImg *gif.GIF, option ResizeOption) {
	// 遍历GIF的每一帧进行缩放
	for i := range gifImg.Image {
		original := gifImg.Image[i]

		// 使用nfnt/resize库对图像进行缩放
		resized := ResizeImage(original, option)

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
}
