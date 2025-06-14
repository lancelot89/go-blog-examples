package pipeline

import (
	"bytes"
	"image"
	"image/jpeg"
)

func Encode(img image.Image, quality int) ([]byte, error) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	return buf.Bytes(), err
}
