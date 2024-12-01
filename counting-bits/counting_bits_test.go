package counting_bits

import (
	"slices"
	"testing"
)

func Test_countBits(t *testing.T) {
	tcases := []struct {
		n      int
		output []int
	}{
		{
			n:      2,
			output: []int{0, 1, 1},
		},
		{
			n:      5,
			output: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tc := range tcases {
		t.Run("countBitsDp", func(t *testing.T) {
			if output := countBitsDp(tc.n); !slices.Equal(output, tc.output) {
				t.Errorf("got %v, expected %v", output, tc.output)
			}
		})

		t.Run("countBitsNaive", func(t *testing.T) {
			if output := countBitsNaive(tc.n); !slices.Equal(output, tc.output) {
				t.Errorf("got %v, expected %v", output, tc.output)
			}
		})
	}
}
