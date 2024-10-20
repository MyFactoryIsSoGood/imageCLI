// imaging/imaging.go
package service

import (
	"image"
)

type Operations struct {
	Resize      *ResizeParams
	Blur        *BlurParams
	Adjust      *AdjustParams
	UseParallel bool
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
