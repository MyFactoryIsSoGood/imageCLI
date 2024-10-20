package service

import (
	"image"
	"image/color"
	"math"
)

// AdjustParams содержит параметры для коррекции изображения.
type AdjustParams struct {
	Saturation float64
	Contrast   float64
	Hue        float64
	Invert     bool
	RedShift   float64
	GreenShift float64
	BlueShift  float64
}

// adjust применяет коррекцию к изображению на основе заданных параметров.
func (is *ImageService) adjust(img image.Image, params AdjustParams) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			rf := float64(r >> 8)
			gf := float64(g >> 8)
			bf := float64(b >> 8)
			af := float64(a >> 8)

			// Инвертирование
			if params.Invert {
				rf = 255 - rf
				gf = 255 - gf
				bf = 255 - bf
			}

			// Сдвиг RGB каналов
			rf = clamp(rf+params.RedShift, 0, 255)
			gf = clamp(gf+params.GreenShift, 0, 255)
			bf = clamp(bf+params.BlueShift, 0, 255)

			// Преобразование в HSL для изменения насыщенности и оттенка
			h, s, l := rgbToHsl(rf/255, gf/255, bf/255)

			// Изменение насыщенности
			if params.Saturation != 0 {
				s += params.Saturation / 100
				s = clamp(s, 0, 1)
			}

			// Изменение оттенка
			if params.Hue != 0 {
				h += params.Hue / 360
				h = math.Mod(h, 1)
				if h < 0 {
					h += 1
				}
			}

			// Преобразование обратно в RGB
			rf, gf, bf = hslToRgb(h, s, l)
			rf *= 255
			gf *= 255
			bf *= 255

			// Изменение контрастности
			if params.Contrast != 0 {
				factor := (259 * (params.Contrast + 255)) / (255 * (259 - params.Contrast))
				rf = clamp(factor*(rf-128)+128, 0, 255)
				gf = clamp(factor*(gf-128)+128, 0, 255)
				bf = clamp(factor*(bf-128)+128, 0, 255)
			}

			dst.Set(x, y, color.NRGBA{
				R: uint8(rf),
				G: uint8(gf),
				B: uint8(bf),
				A: uint8(af),
			})
		}
	}

	return dst
}
