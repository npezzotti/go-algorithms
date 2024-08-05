package zero_one_matrix

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	testCases := map[string]struct {
		matrix [][]int
		want   [][]int
	}{
		"one": {
			matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want:   [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		"two": {
			matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			want:   [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
	}
	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			got := updateMatrix(tc.matrix)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v\n", got, tc.want)
			}
		})
	}
}
