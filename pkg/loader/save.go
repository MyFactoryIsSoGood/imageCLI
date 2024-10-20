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

func SaveImagesToDir(dir string, imgs []ImageFile) error {
	if len(imgs) == 0 {
		return fmt.Errorf("нет изображений для сохранения")
	}

	// Проверяем, является ли путь директорией или файлом
	isSingle := len(imgs) == 1
	ext := ""
	if isSingle {
		ext = strings.ToLower(filepath.Ext(dir))
	}

	if isSingle && ext != "" {
		// Путь предполагается как файл
		return saveImageToFile(dir, imgs[0].Img)
	}

	// Путь предполагается как директория
	// Проверяем существование директории, если нет - создаем
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("не удалось создать директорию '%s': %v", dir, err)
		}
	} else if err != nil {
		return fmt.Errorf("ошибка при проверке директории '%s': %v", dir, err)
	}

	var errs []error

	for i, img := range imgs {
		fileName := img.Name
		outputPath := filepath.Join(dir, fileName)

		// Сохраняем изображение в файл
		err := saveImageToFile(outputPath, img.Img)
		if err != nil {
			errs = append(errs, fmt.Errorf("ошибка при сохранении изображения %d в '%s': %v", i+1, outputPath, err))
		}
	}

	if len(errs) > 0 {
		// Объединяем все ошибки в одну
		errorMessages := "Произошли следующие ошибки при сохранении изображений:\n"
		for _, e := range errs {
			errorMessages += e.Error() + "\n"
		}
		return fmt.Errorf(errorMessages)
	}

	return nil
}
func saveImageToFile(path string, img image.Image) error {
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
