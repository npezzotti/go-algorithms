package subarry_sum_equals_k

import "testing"

func Test_subarraySum(t *testing.T) {
	tcases := []struct {
		nums      []int
		k, output int
	}{
		{
			nums:   []int{1, 1, 1},
			k:      2,
			output: 2,
		},
		{
			nums:   []int{1, 2, 3},
			k:      3,
			output: 2,
		},
		{
			nums:   []int{1, -1, 0},
			k:      0,
			output: 3,
		},
	}

	for _, tt := range tcases {
		t.Run("subarraySumBruteForce", func(t *testing.T) {
			if output := subarraySumBruteForce(tt.nums, tt.k); output != tt.output {
				t.Errorf("got %d for %v, expected %d", output, tt.nums, tt.output)
			}
		})

		t.Run("subarraySum", func(t *testing.T) {
			if output := subarraySum(tt.nums, tt.k); output != tt.output {
				t.Errorf("got %d for %v, expected %d", output, tt.nums, tt.output)
			}
		})
	}
}
