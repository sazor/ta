package ta

import (
	"log"
	"testing"
)

func TestSimpleMA(t *testing.T) {
	input := []float64{25, 85, 65, 12.45, 66.2}
	period := 3
	res, err := SimpleMA(input, period)
	log.Println(res)
	log.Println(err)
	t.Fail()
}

var result []float64

func BenchmarkSimpleMA(b *testing.B) {
	var r []float64
	input := []float64{25, 85, 65, 12.45, 66.2}
	period := 3

	for n := 0; n < b.N; n++ {
		r, _ = SimpleMA(input, period)
	}
	result = r
}
