package imageprocess

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"io"
	"os"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// LoadImageeader 从给定的 io.Reader 中加载指定文件扩展名的图像。
// 参数:
//
//	file: 包含图像数据的 io.Reader。
//	fileExt: 图像文件的扩展名，用于确定图像格式。
//
// 返回值:
//
//	img: 成功加载的图像对象。
//	err: 如果加载失败，返回错误信息。
func DecodeImage(file io.Reader, format Format) (img image.Image, err error) {
	// 根据文件扩展名选择合适的图像解码方法
	switch format {
	case JPEG, JPG:
		// 使用 imaging 包解码 JPG 和 JPEG 格式，并自动处理图像方向
		img, err = imaging.Decode(file, imaging.AutoOrientation(true))
	case PNG:
		// 使用 imaging 包解码 PNG 格式，并自动处理图像方向
		img, err = imaging.Decode(file, imaging.AutoOrientation(true))
	case BMP:
		// 使用 bmp 包解码 BMP 格式
		img, err = bmp.Decode(file)
	case GIF:
		// 使用 gif 包解码 GIF 格式
		img, err = gif.Decode(file)
	case WEBP:
		// 使用 webp 包解码 WEBP 格式
		img, err = webp.Decode(file)
	case TIFF:
		// 使用 tiff 包解码 TIFF 格式
		img, err = tiff.Decode(file)
	default:
		// 如果文件扩展名不在支持的格式列表中，返回错误
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}
	// 如果解码过程中出现错误，返回错误信息
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}
	// 返回成功加载的图像
	return img, nil
}

// EncodeImage 根据指定的格式和质量将图像编码并写入指定的输出流。
// 参数:
//
//	img: 需要编码的图像。
//	w: 写入编码后图像数据的输出流。
//	format: 图像的格式字符串，如"jpeg"、"png"等。
//	quality: 图像的压缩质量，取值范围为1到100。
//
// 返回值:
//
//	如果编码过程中出现错误，返回相应的错误信息。
func EncodeImage(img image.Image, w io.Writer, format Format, quality int) error {
	// 根据文件扩展名获取图像格式。
	f, _ := imaging.FormatFromExtension(string(format))
	// 根据图像格式进行编码。
	switch format {
	case JPEG, JPG:
		// 对于JPEG格式，使用指定的质量进行编码。
		err := imaging.Encode(w, img, f, imaging.JPEGQuality(quality))
		return err
	case PNG:
		// 对于PNG格式，使用最佳压缩级别进行编码。
		err := imaging.Encode(w, img, f, imaging.PNGCompressionLevel(png.BestCompression))
		return err
	case BMP:
		// 对于BMP格式，进行编码。
		err := bmp.Encode(w, img)
		return err
	case GIF:
		// 对于GIF格式，进行编码。
		err := imaging.Encode(w, img, f)
		return err
	case WEBP:
		// 对于WEBP格式，使用指定的质量进行无损编码。
		err := webp.Encode(w, img, &webp.Options{Lossless: false, Quality: float32(quality)})
		return err
	case TIFF:
		// 对于TIFF格式，进行编码。
		err := tiff.Encode(w, img, nil)
		return err
	default:
		// 如果是不支持的图像格式，返回错误信息。
		return fmt.Errorf("unsupported image format: %s", format)
	}
}

func SaveImage(img image.Image, filePath string, format Format, quality int) error {
	// 打开文件，以写模式
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 将图像编码并写入文件
	err = EncodeImage(img, file, format, quality)
	if err != nil {
		return err
	}
	return nil
}
