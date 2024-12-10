package next_permutation

import (
	"slices"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tcases := []struct {
		nums   []int
		output []int
	}{
		{
			nums:   []int{1, 3, 2},
			output: []int{2, 1, 3},
		},
		{
			nums:   []int{3, 2, 1},
			output: []int{1, 2, 3},
		},
		{
			nums:   []int{2, 1, 3},
			output: []int{2, 3, 1},
		},
		{
			nums:   []int{1, 2, 4, 3},
			output: []int{1, 3, 2, 4},
		},
	}

	for _, tc := range tcases {
		nextPermutation(tc.nums)
		if !slices.Equal(tc.nums, tc.output) {
			t.Errorf("got %v, expected %v", tc.nums, tc.output)
		}
	}
}
