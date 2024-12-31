package contiguous_array

import "testing"

func Test_findMaxLength(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{0, 1},
			output: 2,
		},
		{
			nums:   []int{0, 1, 0},
			output: 2,
		},
		{
			nums:   []int{0, 1, 0, 1},
			output: 4,
		},
	}

	for _, tt := range tcases {
		if output := findMaxLength(tt.nums); output != tt.output {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
