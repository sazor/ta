package ta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const float64EqualityThreshold = 1e-2

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

func TestSimpleMA(t *testing.T) {
	type args struct {
		values     []float64
		timePeriod int
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{"valid values", args{[]float64{25, 85, 65, 12.45, 66.2}, 3}, []float64{58.33, 54.15, 47.883}, false},
		{"time period < 1", args{[]float64{25, 85, 65, 12.45, 66.2}, 0}, nil, true},
		{"time period > length of values", args{[]float64{25, 85, 65, 12.45, 66.2}, 6}, nil, true},
		{"time period = length of values", args{[]float64{12, 13, 15, 20}, 4}, []float64{15}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SimpleMA(tt.args.values, tt.args.timePeriod)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.InDeltaSlicef(t, tt.want, got, float64EqualityThreshold,
				"expected: %v, got: %v", tt.want, got)
		})
	}
}
