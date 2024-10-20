package service

import (
	"image"
)

// ResizeParams - параметры изменения размера
type ResizeParams struct {
	Width  int
	Height int
}

// resize изменяет размер изображения с помощью метода ближайшего соседа
func (is *ImageService) resize(img image.Image, width, height int) image.Image {
	srcBounds := img.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		srcY := y * srcBounds.Dy() / height
		for x := 0; x < width; x++ {
			srcX := x * srcBounds.Dx() / width
			color := img.At(srcX+srcBounds.Min.X, srcY+srcBounds.Min.Y)
			dst.Set(x, y, color)
		}
	}

	return dst
}
