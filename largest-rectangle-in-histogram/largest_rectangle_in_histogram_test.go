package largest_rectangle_in_histogram

import (
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	tcases := []struct {
		heights []int
		output  int
	}{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			output:  10,
		},
		{
			heights: []int{2, 4},
			output:  4,
		},
		{
			heights: []int{2, 1, 2},
			output:  3,
		},
	}

	for _, tc := range tcases {
		if output := largestRectangleArea(tc.heights); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
