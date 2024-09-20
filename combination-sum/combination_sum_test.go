package combination_sum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tcases := []struct {
		candidates []int
		target     int
		output     [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			output:     [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			output:     [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			output:     [][]int{},
		},
	}

	for _, tc := range tcases {
		if output := combinationSum(tc.candidates, tc.target); !reflect.DeepEqual(output, tc.output) {
			t.Fatalf("got %v, want %v", output, tc.output)
		}
	}
}
