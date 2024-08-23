package imageprocess

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

func AdjustBlur(img image.Image, option BlurOption) image.Image {
	return imaging.Blur(img, float64(option.Radius))
}

func AddjustSharpen(img image.Image, option SharpenOption) image.Image {
	return imaging.Sharpen(img, float64(option.Value/100))
}
func AddjustBright(img image.Image, option BrightnessOption) image.Image {
	return imaging.AdjustBrightness(img, float64(option.Value))
}

func AddjustContrast(img image.Image, option ContrastOption) image.Image {
	return imaging.AdjustContrast(img, float64(option.Value))
}

func AddjustGamma(img image.Image, option GammaOption) image.Image {
	return imaging.AdjustGamma(img, float64(option.Value/100))
}

func AddjustSaturation(img image.Image, option SaturationOption) image.Image {
	return imaging.AdjustSaturation(img, float64(option.Value))
}

func AddjustRotate(img image.Image, option RotateOption) image.Image {
	return imaging.Rotate(img, -float64(option.Value), color.RGBA64{255, 255, 255, 0})
}
