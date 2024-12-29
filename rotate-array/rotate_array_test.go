package rotate_array

import (
	"slices"
	"testing"
)

func Test_rotate(t *testing.T) {
	tcases := []struct {
		nums   []int
		k      int
		output []int
	}{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			k:      3,
			output: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:   []int{-1, -100, 3, 99},
			k:      2,
			output: []int{3, 99, -1, -100},
		},
		{
			nums:   []int{-1},
			k:      2,
			output: []int{-1},
		},
		{
			nums:   []int{1, 2, 3},
			k:      4,
			output: []int{3, 1, 2},
		},
		{
			nums:   []int{1, 2},
			k:      5,
			output: []int{2, 1},
		},
	}

	for _, tt := range tcases {
		if rotate(tt.nums, tt.k); !slices.Equal(tt.nums, tt.output) {
			t.Errorf("got %v, expected %v", tt.nums, tt.output)
		}
	}
}
