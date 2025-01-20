package pow

import (
	"strconv"
	"testing"
)

func Test_myPow(t *testing.T) {
	tcases := []struct {
		x      float64
		n      int
		output float64
	}{
		{
			x:      2.00000,
			n:      10,
			output: 1024.00000,
		},
		{
			x:      2.10000,
			n:      3,
			output: 9.26100,
		},
		{
			x:      2.00000,
			n:      -2,
			output: 0.25000,
		},
	}

	for _, tt := range tcases {
		if output := myPow(tt.x, tt.n); strconv.FormatFloat(output, 'f', 5, 64) != strconv.FormatFloat(tt.output, 'f', 5, 64) {
			t.Errorf("got %f, expected %f", output, tt.output)
		}
	}
}
