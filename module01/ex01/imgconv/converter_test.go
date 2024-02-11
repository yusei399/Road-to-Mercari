package imgconv

import (
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func setupTestDir(t *testing.T) string {
	t.Helper()

	dir, err := ioutil.TempDir("", "imgconv_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	sampleJPGPath := filepath.Join(dir, "sample.jpg")
	jpgFile, err := os.Create(sampleJPGPath)
	if err != nil {
		t.Fatalf("Failed to create sample JPG file: %v", err)
	}
	defer jpgFile.Close()

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}
	if err := jpeg.Encode(jpgFile, img, nil); err != nil {
		t.Fatalf("Failed to write sample JPG content: %v", err)
	}

	return dir
}

func teardownTestDir(t *testing.T, dir string) {
	t.Helper()
	if err := os.RemoveAll(dir); err != nil {
		t.Errorf("Failed to remove test dir: %v", err)
	}
}

func TestConvertJPGToPNG(t *testing.T) {
	t.Parallel()

	dir := setupTestDir(t)
	defer teardownTestDir(t, dir)

	converter := NewConverter()
	if err := converter.ConvertJPGToPNG(dir); err != nil {
		t.Errorf("ConvertJPGToPNG failed: %v", err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("Failed to read test dir: %v", err)
	}

	foundPNG := false
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".png") {
			foundPNG = true
			break
		}
	}

	if !foundPNG {
		t.Errorf("No PNG file found after conversion")
	}
}

