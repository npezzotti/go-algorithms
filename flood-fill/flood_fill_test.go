package flood_fill

import "testing"

func TestFloodFill(t *testing.T) {
	got := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 0, 1},
	}

	want := [][]int{
		{2, 2, 2},
		{0, 2, 0},
		{0, 0, 1},
	}

	floodFill(got, 1, 1, 2)

	for rowIdx := range got {
		for colIdx, val := range got[rowIdx] {
			if want[rowIdx][colIdx] != val {
				t.Fatalf("got %+v, want %+v\n", got, want)
			}
		}
	}
}
