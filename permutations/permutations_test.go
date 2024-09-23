package permutations

import (
	"reflect"
	"testing"
)

func TestPermutate(t *testing.T) {
	tcases := []struct {
		nums   []int
		output [][]int
	}{
		{
			nums: []int{1, 2, 3},
			output: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2}},
		},
		{
			nums:   []int{0, 1},
			output: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums:   []int{1},
			output: [][]int{{1}},
		},
	}

	for _, tc := range tcases {
		if output := permute(tc.nums); !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("got %v, want %v", output, tc.output)
		}
	}
}
