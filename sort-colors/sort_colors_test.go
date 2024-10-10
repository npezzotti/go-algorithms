package sort_colors

import (
	"slices"
	"testing"
)

func TestSortColors(t *testing.T) {
	tcases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{2, 0, 2, 1, 1, 0},
			output: []int{0, 0, 1, 1, 2, 2},
		},
		{
			input:  []int{2, 0, 1},
			output: []int{0, 1, 2},
		},
	}

	for _, tc := range tcases {
		if sortColors(tc.input); !slices.Equal(tc.input, tc.output) {
			t.Fatalf("expected %v, got %v", tc.output, tc.input)
		}
	}
}
