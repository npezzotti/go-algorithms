package search_2d_matrix

import "testing"

func Test_searchMatrix(t *testing.T) {
	tcases := []struct {
		matrix [][]int
		target int
		output bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			output: true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			output: false,
		},
		{
			matrix: [][]int{{1}},
			target: 1,
			output: true,
		},
		{
			matrix: [][]int{{1}},
			target: 2,
			output: false,
		},
		{
			matrix: [][]int{{1}, {3}, {5}},
			target: 3,
			output: true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target: 11,
			output: true,
		},
	}

	for _, tt := range tcases {
		if output := searchMatrix(tt.matrix, tt.target); output != tt.output {
			t.Errorf("got %t for matrix %v and target %d, expected %t", output, tt.matrix, tt.target, tt.output)
		}
	}
}
