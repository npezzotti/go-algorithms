package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	tcases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
	}

	for _, tc := range tcases {
		if idx := search(tc.nums, tc.target); idx != tc.want {
			t.Fatalf("got %d, want %d", idx, tc.want)
		}
	}
}
