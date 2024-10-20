package service

import (
	"image"
	"imageCLI/pkg/loader"
	"runtime"
	"sync"
)

type ImageService struct {
	maxGoroutines int
}

func NewImageService(maxGoroutines int) *ImageService {
	if maxGoroutines <= 0 {
		maxGoroutines = runtime.NumCPU()
	}

	iService := ImageService{
		maxGoroutines: maxGoroutines,
	}

	return &iService
}

func (is *ImageService) ProcessImages(images []loader.ImageFile, operations Operations) ([]loader.ImageFile, error) {
	var err error

	if operations.UseParallel {
		var wg sync.WaitGroup
		errChan := make(chan error, 1)

		for i, img := range images {
			wg.Add(1)
			go func(i int, img image.Image) {
				defer wg.Done()
				processedImg, procErr := is.process(img, operations)
				if procErr != nil {
					select {
					case errChan <- procErr:
					default:
					}
					return
				}
				images[i].Img = processedImg
			}(i, img.Img)
		}

		wg.Wait()

		select {
		case err = <-errChan:
			return nil, err
		default:
		}
	} else {
		for i, img := range images {
			processedImg, procErr := is.process(img.Img, operations)
			if procErr != nil {
				return nil, procErr
			}
			images[i].Img = processedImg
		}
	}

	return images, nil
}
