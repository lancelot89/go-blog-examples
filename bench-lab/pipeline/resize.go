package pipeline

import (
	"image"

	"golang.org/x/image/draw"
)

func Resize(img image.Image, w, h int) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.ApproxBiLinear.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)
	return dst
}
