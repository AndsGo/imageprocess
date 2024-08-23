package imageprocess

import (
	"image"

	"github.com/AndsGo/imageprocess/fonts"
	"github.com/fogleman/gg"
)

func WarterMarkText(img image.Image, option TextWatermarkOption) image.Image {
	dc := gg.NewContextForImage(img)
	// 设置字体大小和颜色
	if option.Color != nil {
		dc.SetRGBA(float64(option.Color.R)/255, float64(option.Color.G)/255, float64(option.Color.B)/255, float64(option.Opacity)/100) // 白色，透明度60%
	}
	dc.SetFontFace(fonts.GetFont(fonts.ZH_CN, option.Size))
	// 在图片中心绘制文字
	text := option.Text
	// 获取文字位置
	w, h := dc.MeasureString(text)
	x, y := TextWatermarkPositionPoint(img, option, w, h)
	dc.DrawStringAnchored(text, x, y, 0, 0)
	// 旋转
	dc.Rotate(float64(option.Rotate))
	// voffset
	dc.SetDashOffset(float64(option.Voffset))
	return dc.Image()
}
