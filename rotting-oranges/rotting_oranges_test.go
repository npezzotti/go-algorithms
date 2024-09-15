package rotting_oranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	tcases := []struct {
		grid    [][]int
		minutes int
	}{
		{
			grid:    [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			minutes: 4,
		},
		{
			grid:    [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			minutes: -1,
		},
		{
			grid:    [][]int{{0, 2}},
			minutes: 0,
		},
	}

	for _, tc := range tcases {
		if minutes := orangesRotting(tc.grid); minutes != tc.minutes {
			t.Fatalf("expected %d, got %d", tc.minutes, minutes)
		}
	}
}
