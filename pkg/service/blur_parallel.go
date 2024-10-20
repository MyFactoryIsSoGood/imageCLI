package service

import (
	"image"
	"sync"
)

// parallelBlur - разбиваем изображение на чанки и блюрим
func (is *ImageService) parallelBlur(img image.Image, radius float64) (image.Image, error) {
	if radius <= 0 {
		return img, nil
	}

	numThreads := is.maxGoroutines
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	kernelRadius := int(radius*3 + 0.5)

	blockHeights := make([][2]int, numThreads)
	step := height / numThreads
	for i := 0; i < numThreads; i++ {
		startY := i * step
		endY := startY + step
		if i == numThreads-1 {
			endY = height
		}

		startYOverlap := startY
		if startYOverlap > 0 {
			startYOverlap -= kernelRadius
		}
		endYOverlap := endY + kernelRadius
		if endYOverlap > height {
			endYOverlap = height
		}
		blockHeights[i] = [2]int{startYOverlap, endYOverlap}
	}

	var wg sync.WaitGroup
	blocks := make([]*image.RGBA, numThreads)
	imgRGBA := toRGBA(img)

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			startY := blockHeights[i][0]
			endY := blockHeights[i][1]
			subImg := imgRGBA.SubImage(image.Rect(0, startY, width, endY)).(*image.RGBA)

			blurredSubImg := is.blur(subImg, radius).(*image.RGBA)

			cropStartY := startY + kernelRadius
			if i == 0 {
				cropStartY = 0
			}
			cropEndY := endY - kernelRadius
			if i == numThreads-1 {
				cropEndY = height
			}

			croppedImg := blurredSubImg.SubImage(image.Rect(
				0,
				cropStartY,
				width,
				cropEndY,
			)).(*image.RGBA)

			blocks[i] = croppedImg
		}(i)
	}

	wg.Wait()

	finalImg := image.NewRGBA(image.Rect(0, 0, width, height))
	currentY := 0
	for _, block := range blocks {
		drawAt(finalImg, block)
		currentY += block.Bounds().Dy()
	}

	return finalImg, nil
}
