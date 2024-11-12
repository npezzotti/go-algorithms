package mininum_height_trees

import (
	"slices"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	tcases := []struct {
		n      int
		edges  [][]int
		output []int
	}{
		{
			n:      4,
			edges:  [][]int{{1, 0}, {1, 2}, {1, 3}},
			output: []int{1},
		},
		{
			n:      6,
			edges:  [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			output: []int{3, 4},
		},
		{
			n:      1,
			edges:  [][]int{},
			output: []int{0},
		},
	}

	for _, tc := range tcases {
		if output := findMinHeightTrees(tc.n, tc.edges); !slices.Equal(output, tc.output) {
			t.Errorf("expected %+v, got %+v", tc.output, output)
		}
	}
}
