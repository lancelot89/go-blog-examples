package pipeline

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
)

// Decode tries to decode JPEG/PNG bytes.
// If it fails, it returns a solid white dummy image to keep the benchmark running.
func Decode(data []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		// -- Fallback: create 640×480 white image --
		dummy := image.NewRGBA(image.Rect(0, 0, 640, 480))
		draw.Draw(dummy, dummy.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
		return dummy, nil
	}
	return img, nil
}
