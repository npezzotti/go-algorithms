package pacific_atlantic_water_flow

import (
	"reflect"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	tcases := []struct {
		heights [][]int
		output  [][]int
	}{
		{
			heights: [][]int{
				{13, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4}},
			output: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			heights: [][]int{{1}},
			output:  [][]int{{0, 0}},
		},
	}

	for _, tc := range tcases {
		if output := pacificAtlantic(tc.heights); !reflect.DeepEqual(output, tc.output) {
			t.Errorf("got %v, expected %v", output, tc.output)
		}
	}
}
