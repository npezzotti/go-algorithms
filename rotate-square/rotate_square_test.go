package rotate_square

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	tcases := []struct {
		matrix [][]int
		output [][]int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			output: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tt := range tcases {
		if rotate(tt.matrix); !reflect.DeepEqual(tt.matrix, tt.output) {
			t.Errorf("got %v, expected %v", tt.matrix, tt.output)
		}
	}
}
