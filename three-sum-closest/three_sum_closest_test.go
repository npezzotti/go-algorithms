package three_sum_closest

import "testing"

func Test_threeSumClosest(t *testing.T) {
	tcases := []struct {
		nums           []int
		target, output int
	}{
		{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			output: 2,
		},
		{
			nums:   []int{0, 0, 0},
			target: 1,
			output: 0,
		},
		{
			nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
			output: -2,
		},
		{
			nums:   []int{2, 5, 6, 7},
			target: 16,
			output: 15,
		},
	}

	for _, tt := range tcases {
		if output := threeSumClosest(tt.nums, tt.target); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
