package imageprocess

import (
	"errors"
	"strings"
)

const (
	JPG  = "jpg"
	JPEG = "jpeg"
	PNG  = "png"
	BMP  = "bmp"
	GIF  = "gif"
	WEBP = "webp"
	TIFF = "tiff"
)

// Supported formats
var supportedFormats = map[string]Format{
	"jpg":  JPG,
	"jpeg": JPEG,
	"png":  PNG,
	"gif":  GIF,
	"tif":  TIFF,
	"tiff": TIFF,
	"bmp":  BMP,
	"webp": WEBP,
}

// Format is an image file format.
type Format string

// ErrUnsupportedFormat means the given image format is not supported.
var ErrUnsupportedFormat = errors.New("imaging: unsupported image format")
var ErrUnsupportedParameter = errors.New("imaging: unsupported image parameter")

// FormatFromExtension parses image format from filename extension:
func FormatFromExtension(ext string) (Format, error) {
	extArr := strings.Split(strings.ToLower(ext), ".")
	if f, ok := supportedFormats[extArr[len(extArr)-1]]; ok {
		return f, nil
	}
	return "", ErrUnsupportedFormat
}

func FormatFromFilename(filename string) (Format, error) {
	return FormatFromExtension(strings.ToLower(strings.TrimPrefix(filename, ".")))
}
