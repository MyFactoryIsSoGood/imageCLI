package service

import (
	"image"
	"image/color"
	"math"
)

// BlurParams содержит параметры для размытия.
type BlurParams struct {
	Radius float64
}

// blur применяет гауссово размытие к изображению с заданным радиусом.
func (is *ImageService) blur(img image.Image, radius float64) image.Image {
	if radius <= 0 {
		return img
	}

	// Создаем ядро размытия
	kernelSize := int(math.Ceil(radius)*2 + 1)
	kernel := make([]float64, kernelSize)
	sigma := radius / 3
	sum := 0.0

	for i := 0; i < kernelSize; i++ {
		x := float64(i - kernelSize/2)
		value := math.Exp(-(x * x) / (2 * sigma * sigma))
		kernel[i] = value
		sum += value
	}

	// Нормализуем ядро
	for i := range kernel {
		kernel[i] /= sum
	}

	// Применяем размытие по горизонтали и вертикали
	tempImg := convolve(img, kernel, true)
	blurredImg := convolve(tempImg, kernel, false)

	return blurredImg
}

// convolve выполняет свертку изображения с ядром
func convolve(img image.Image, kernel []float64, horizontal bool) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	kernelSize := len(kernel)
	radius := kernelSize / 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var r, g, b, a float64
			for k := -radius; k <= radius; k++ {
				weight := kernel[k+radius]
				var srcX, srcY int
				if horizontal {
					srcX = x + k
					srcY = y
				} else {
					srcX = x
					srcY = y + k
				}
				// Проверяем границы
				if srcX < bounds.Min.X {
					srcX = bounds.Min.X
				}
				if srcX >= bounds.Max.X {
					srcX = bounds.Max.X - 1
				}
				if srcY < bounds.Min.Y {
					srcY = bounds.Min.Y
				}
				if srcY >= bounds.Max.Y {
					srcY = bounds.Max.Y - 1
				}
				srcR, srcG, srcB, srcA := img.At(srcX, srcY).RGBA()
				r += float64(srcR>>8) * weight
				g += float64(srcG>>8) * weight
				b += float64(srcB>>8) * weight
				a += float64(srcA>>8) * weight
			}
			dst.Set(x, y, color.NRGBA{
				R: uint8(clamp(r, 0, 255)),
				G: uint8(clamp(g, 0, 255)),
				B: uint8(clamp(b, 0, 255)),
				A: uint8(clamp(a, 0, 255)),
			})
		}
	}

	return dst
}
