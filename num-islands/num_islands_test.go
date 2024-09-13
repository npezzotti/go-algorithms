package num_islands

import "testing"

func TestNumIslands(t *testing.T) {
	tcases := []struct {
		grid [][]byte
		want int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}

	for _, tc := range tcases {
		if got := numIslands(tc.grid); got != tc.want {
			t.Fatalf("got %d, want %d", got, tc.want)
		}
	}
}
