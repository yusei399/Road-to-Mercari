//JPGファイルをPNGファイルに変換する機能
package converter

import (
	"image"
	_ "image/jpeg" // JPEG形式の画像をデコードするために必要
	"image/png"    // PNG形式でエンコードするために必要
	"os"
	"path/filepath"
)

//画像の形式変換を行うための型
type Converter struct{}

// NewConverter は新しいConverterのインスタンスを生成
func NewConverter() *Converter {
	return &Converter{}
}

// ConvertJPGToPNG は指定されたディレクトリ内の全てのJPGファイルをPNGに変換
func (c *Converter) ConvertJPGToPNG(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".jpg" {
			return convertToPNG(path)
		}
		return nil
	})
}

// JPGファイルをPNGファイルに変換
func convertToPNG(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + ".png"
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

