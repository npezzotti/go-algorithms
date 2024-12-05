package squares_of_a_sorted_array

import (
	"slices"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tcases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			input:  []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{0, 2},
			output: []int{0, 4},
		},
		{
			input:  []int{-5, -3, -2, -1},
			output: []int{1, 4, 9, 25},
		},
		{
			input:  []int{-5, -3, -2, 1},
			output: []int{1, 4, 9, 25},
		},
	}

	for _, tc := range tcases {
		if output := sortedSquares(tc.input); !slices.Equal(output, tc.output) {
			t.Errorf("got %+v, expected %+v", output, tc.output)
		}
	}
}
