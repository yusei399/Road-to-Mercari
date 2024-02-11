// Package converter converts JPG files to PNG files.
package converter

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import for JPG decoding support
	"image/png"    // Import for PNG encoding
	"os"
	"path/filepath"
)

// Converter is a type that performs image format conversions.
type Converter struct{}

// NewConverter creates and returns a new instance of Converter.
func NewConverter() *Converter {
	return &Converter{}
}

// ConvertJPGToPNG converts all JPG files in the specified directory to PNG files.
func (c *Converter) ConvertJPGToPNG(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking through directory: %v", err)
		}
		if !info.IsDir() && filepath.Ext(path) == ".jpg" {
			return convertToPNG(path)
		}
		return nil
	})
}

// convertToPNG converts a single JPG file to a PNG file.
func convertToPNG(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", filePath, err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding JPG file %s: %v", filePath, err)
	}

	outputPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + ".png"
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating PNG file %s: %v", outputPath, err)
	}
	defer outFile.Close()

	if err := png.Encode(outFile, img); err != nil {
		return fmt.Errorf("error encoding file to PNG %s: %v", outputPath, err)
	}

	return nil
}

