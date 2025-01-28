package set_matrix_zeroes

import (
	"reflect"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	tcases := []struct {
		matrix, output [][]int
	}{
		{
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			output: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			output: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			matrix: [][]int{{1, 0, 3}},
			output: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tcases {
		setZeroes(tt.matrix)
		if !reflect.DeepEqual(tt.matrix, tt.output) {
			t.Errorf("got %v, expected %v", tt.matrix, tt.output)
		}
	}
}
