// loader/load.go
package loader

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
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

// LoadImageFromFile загружает изображение из файла.
func loadImageFromFile(path string) (ImageFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return ImageFile{}, err
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(path))
	var img image.Image
	switch ext {
	case ".png":
		img, err = png.Decode(file)
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	default:
		err = fmt.Errorf("неподдерживаемый формат файла: %s", ext)
	}

	imageFile := ImageFile{
		Name: filepath.Base(path),
		Img:  img,
	}

	return imageFile, err
}

// Возвращает срез загруженных ImageFile и ошибку, если таковая возникла.
func loadImagesFromDirectory(inputDir string, parallel bool) ([]ImageFile, error) {
	var images []ImageFile

	files, err := getFiles(inputDir)
	if err != nil {
		return nil, err
	}

	if parallel {
		semaphore := make(chan struct{}, runtime.NumCPU())
		var wg sync.WaitGroup
		var mu sync.Mutex // Для безопасного добавления в срез images
		var firstErr error

		for _, filePath := range files {
			wg.Add(1)
			semaphore <- struct{}{} // Захватываем слот

			go func(path string) {
				defer wg.Done()
				defer func() { <-semaphore }() // Освобождаем слот

				img, err := loadImageFromFile(path)
				if err != nil {
					// Сохраняем первую возникшую ошибку
					mu.Lock()
					if firstErr == nil {
						firstErr = err
					}
					mu.Unlock()
					return
				}

				// Добавляем изображение в срез
				mu.Lock()
				images = append(images, img)
				mu.Unlock()
			}(filePath)
		}

		wg.Wait()

		if firstErr != nil {
			return images, firstErr
		}

	} else {
		// Последовательная загрузка
		for _, filePath := range files {
			img, err := loadImageFromFile(filePath)
			if err != nil {
				return images, err
			}
			images = append(images, img)
		}
	}

	return images, nil
}

// getFiles возвращает срез всех файлов в указанной директории с поддерживаемыми расширениями.
func getFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			if isSupportedFormat(ext) {
				files = append(files, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("ошибка при обходе директории '%s': %v", dir, err)
	}

	return files, nil
}

// isSupportedFormat проверяет, поддерживается ли формат файла.
func isSupportedFormat(ext string) bool {
	switch ext {
	case ".png", ".jpg", ".jpeg":
		return true
	default:
		return false
	}
}
