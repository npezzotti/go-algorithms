package kth_largest_element_in_an_array

import "testing"

func Test_findKthLargest(t *testing.T) {
	tcases := []struct {
		nums      []int
		k, output int
	}{
		{
			nums:   []int{3, 2, 1, 5, 6, 4},
			k:      2,
			output: 5,
		},
		{
			nums:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			output: 4,
		},
	}

	for _, tt := range tcases {
		if output := findKthLargest(tt.nums, tt.k); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
