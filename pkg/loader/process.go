// loader/process.go
package loader

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
)

type ImageFile struct {
	Name string
	Img  image.Image
}

func LoadImages(path string, parallel bool) ([]ImageFile, error) {
	var images []ImageFile

	inputInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if inputInfo.IsDir() {
		images, err = loadImagesFromDirectory(path, parallel)
	} else {
		img, err := loadImageFromFile(path)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}

func SaveImagesToDir(dir string, imgs []ImageFile) error {
	fmt.Println(dir)
	// Проверяем существование директории, если нет - создаем
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("не удалось создать директорию '%s': %v", dir, err)
		}
	}

	var errs []error

	for i, img := range imgs {
		fileName := img.Name // Здесь используем PNG формат по умолчанию
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
