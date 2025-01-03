package find_k_closest_elements

import (
	"slices"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	tcases := []struct {
		arr    []int
		k, x   int
		output []int
	}{
		{
			arr:    []int{1, 2, 3, 4, 5},
			k:      4,
			x:      3,
			output: []int{1, 2, 3, 4},
		},
		{
			arr:    []int{1, 1, 2, 3, 4, 5},
			k:      4,
			x:      -1,
			output: []int{1, 1, 2, 3},
		},
		{
			arr:    []int{1},
			k:      1,
			x:      1,
			output: []int{1},
		},
	}

	for _, tt := range tcases {
		if output := findClosestElements(tt.arr, tt.k, tt.x); !slices.Equal(output, tt.output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
