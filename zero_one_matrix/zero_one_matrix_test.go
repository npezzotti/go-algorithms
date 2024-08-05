package zero_one_matrix

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	got := updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	want := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v\n", got, want)
	}
}
