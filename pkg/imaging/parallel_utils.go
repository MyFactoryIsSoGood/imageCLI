// imaging/parallel_utils.go
package imaging

import (
	"image"
)

// toRGBA преобразует изображение в формат RGBA.
func toRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbaImg.Set(x, y, img.At(x, y))
		}
	}
	return rgbaImg
}

// drawAt копирует src в dst
func drawAt(dst *image.RGBA, src *image.RGBA) {
	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dst.Set(x, y, src.At(x, y))
		}
	}
}
