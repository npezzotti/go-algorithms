package combination_sum_4

import "testing"

func Test_combinationSum4(t *testing.T) {
	tcases := []struct {
		nums           []int
		target, output int
	}{
		{
			nums:   []int{1, 2, 3},
			target: 4,
			output: 7,
		},
		{
			nums:   []int{9},
			target: 3,
			output: 0,
		},
	}

	for _, tt := range tcases {
		if output := combinationSum4(tt.nums, tt.target); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
