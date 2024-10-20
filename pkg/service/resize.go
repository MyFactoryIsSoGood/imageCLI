// imaging/resize.go
package service

import (
	"image"
)

// ResizeParams contains parameters for resizing.
type ResizeParams struct {
	Width  int
	Height int
}

// resize resizes the image to the specified width and height.
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
