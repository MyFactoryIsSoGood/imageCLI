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

type Operations struct {
	Resize      *ResizeParams
	Blur        *BlurParams
	Adjust      *AdjustParams
	UseParallel bool
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

func (is *ImageService) process(img image.Image, ops Operations) (image.Image, error) {
	var err error

	if ops.Resize != nil {
		img = is.resize(img, ops.Resize.Width, ops.Resize.Height)
	}

	if ops.Blur != nil {
		if ops.UseParallel {
			img, err = is.parallelBlur(img, ops.Blur.Radius)
			if err != nil {
				return nil, err
			}
		} else {
			img = is.blur(img, ops.Blur.Radius)
		}
	}

	if ops.Adjust != nil {
		if ops.UseParallel {
			img, err = is.parallelAdjust(img, *ops.Adjust)
			if err != nil {
				return nil, err
			}
		} else {
			img = is.adjust(img, *ops.Adjust)
		}
	}

	return img, nil
}
