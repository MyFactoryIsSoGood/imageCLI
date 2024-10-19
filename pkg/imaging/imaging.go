// imaging/imaging.go
package imaging

import (
	"image"
)

type Operations struct {
	Resize      *ResizeParams
	Blur        *BlurParams
	Adjust      *AdjustParams
	UseParallel bool
}

func Process(img image.Image, ops Operations) (image.Image, error) {
	var err error

	if ops.Resize != nil {
		img = Resize(img, ops.Resize.Width, ops.Resize.Height)
	}

	if ops.Blur != nil {
		if ops.UseParallel {
			img, err = ParallelBlur(img, ops.Blur.Radius)
			if err != nil {
				return nil, err
			}
		} else {
			img = Blur(img, ops.Blur.Radius)
		}
	}

	if ops.Adjust != nil {
		if ops.UseParallel {
			img, err = ParallelAdjust(img, *ops.Adjust)
			if err != nil {
				return nil, err
			}
		} else {
			img = Adjust(img, *ops.Adjust)
		}
	}

	return img, nil
}
