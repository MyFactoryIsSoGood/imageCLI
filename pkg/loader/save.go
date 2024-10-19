// loader/save.go
package loader

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func saveImageToFile(path string, img image.Image) error {
	fmt.Println(path)
	ext := strings.ToLower(filepath.Ext(path))
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("не удалось создать директорию '%s': %v", dir, err)
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("не удалось создать файл '%s': %v", path, err)
	}
	defer file.Close()

	switch ext {
	case ".png":
		err = png.Encode(file, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 90})
	default:
		return fmt.Errorf("неподдерживаемый формат файла: %s", ext)
	}

	if err != nil {
		return fmt.Errorf("ошибка при кодировании изображения в '%s': %v", path, err)
	}

	return nil
}
