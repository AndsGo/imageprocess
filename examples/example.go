package main

import (
	"fmt"
	"image/gif"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/AndsGo/imageprocess"
)

// 文件夹，有need change it
var fileFolders = "./"

func main() {
	http.HandleFunc("/file/", fileHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// 获取文件名称
	fileName := strings.TrimPrefix(r.URL.Path, "/file/")
	// 打开文件
	file, err := os.Open(fmt.Sprintf("%s%s", fileFolders, fileName))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	// 获取参数
	// 获取文件后缀
	f, err := imageprocess.FormatFromExtension(fileName)
	if err != nil {
		// 将处理后的文件内容写入响应
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	//处理处理参数
	ossParams := r.URL.Query().Get("x-oss-process")
	if ossParams == "" {
		//无需处理
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	options, err := imageprocess.ParseOptions(ossParams)
	if err != nil {
		http.Error(w, fmt.Sprintf("ParseOptions %s", err.Error()), http.StatusInternalServerError)
		return
	}
	if len(options) == 0 {
		//无需处理
		if _, err := io.Copy(w, file); err != nil {
			http.Error(w, "Failed to send file", http.StatusInternalServerError)
		}
		return
	}
	//处理图片
	err = processImg(file, w, f, options)
	if err != nil {
		http.Error(w, fmt.Sprintf("processFile %s", err.Error()), http.StatusInternalServerError)
	}
}

// 进行转换
func processImg(file io.Reader, w io.Writer, f imageprocess.Format, options []imageprocess.Option) error {
	if f == imageprocess.GIF {
		imgGif, err := gif.DecodeAll(file)
		if err != nil {
			return err
		}
		return imageprocess.ProcessGif(imgGif, w, options)
	} else {
		img, err := imageprocess.DecodeImage(file, f)
		if err != nil {
			return err
		}
		return imageprocess.Process(img, w, f, options)
	}
}
