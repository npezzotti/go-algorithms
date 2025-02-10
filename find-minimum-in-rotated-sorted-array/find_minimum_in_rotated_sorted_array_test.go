package find_minimum_in_rotated_sorted_array

import "testing"

func Test_findMin(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{3, 4, 5, 1, 2},
			output: 1,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			output: 0,
		},
		{
			nums:   []int{11, 13, 15, 17},
			output: 11,
		},
	}

	for _, tt := range tcases {
		if output := findMin(tt.nums); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
