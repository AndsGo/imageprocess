package imageprocess

import (
	"image"

	"github.com/disintegration/imaging"
)

func CropImage(img image.Image, option CropOption) image.Image {
	// 获取img最大值
	if option.Position == "" && option.X == 0 && option.Y == 0 && option.Width == 0 && option.Height == 0 {
		// 无需裁剪
		return img
	}
	if option.Width == 0 {
		option.Width = img.Bounds().Dx()
	}
	if option.Height == 0 {
		option.Height = img.Bounds().Dy()
	}

	if option.Position != "" {
		x, y := CropPositionPoint(img, option.Position, float64(option.Width), float64(option.Height))
		return imaging.Crop(img, image.Rect(int(x), int(y), int(x)+option.Width, int(y)+option.Height))
	}
	return imaging.Crop(img, image.Rect(option.X, option.Y, option.X+option.Width, option.Y+option.Height))
}
