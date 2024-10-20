package service

import (
	"image"
	"sync"
)

func (is *ImageService) parallelAdjust(img image.Image, params AdjustParams) (image.Image, error) {
	numThreads := is.maxGoroutines
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	step := height / numThreads

	var wg sync.WaitGroup
	blocks := make([]*image.RGBA, numThreads)
	imgRGBA := toRGBA(img)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			startY := i * step
			endY := startY + step
			if i == numThreads-1 {
				endY = height
			}
			subImg := imgRGBA.SubImage(image.Rect(0, startY, width, endY)).(*image.RGBA)

			adjustedSubImg := is.adjust(subImg, params).(*image.RGBA)

			blocks[i] = adjustedSubImg
		}(i)
	}

	wg.Wait()

	finalImg := image.NewRGBA(bounds)
	currentY := 0
	for _, block := range blocks {
		drawAt(finalImg, block)
		currentY += block.Bounds().Dy()
	}

	return finalImg, nil
}
