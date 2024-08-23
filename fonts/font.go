package fonts

import (
	"embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.ttf
var fonts embed.FS

const (
	ZH_CN = "zh_cn.ttf"
)

func GetFont(f string, size int) font.Face {
	if f == "" {
		f = ZH_CN
	}
	fontBytes, _ := fonts.ReadFile(f)
	fnt, _ := opentype.Parse(fontBytes)
	face, _ := opentype.NewFace(fnt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
	return face
}
