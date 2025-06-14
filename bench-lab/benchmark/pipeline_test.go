package benchmark

import (
	_ "image/jpeg"
	"testing"

	_ "embed"

	"github.com/izayo/go-blog-examples/bench-lab/pipeline"
)

//go:embed image-data/sample.jpg
var sampleImage []byte

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img, _ := pipeline.Decode(sampleImage)
		img = pipeline.Resize(img, 320, 240)
		_, _ = pipeline.Encode(img, 80)
	}
}
