package spiral_matrix

import (
	"slices"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tcases := []struct {
		input  [][]int
		output []int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			output: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			output: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
	}

	for _, tc := range tcases {
		if output := spiralOrder(tc.input); !slices.Equal(output, tc.output) {
			t.Errorf("expected %v, got %v", tc.output, output)
		}
	}
}
