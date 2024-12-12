package maximum_product_subarray

import "testing"

func Test_maxProduct(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{2, 3, -2, 4},
			output: 6,
		},
		{
			nums:   []int{-2, 0, -1},
			output: 0,
		},
		{
			nums:   []int{-4, -3},
			output: 12,
		},
		{
			nums:   []int{-2},
			output: -2,
		},
		{
			nums:   []int{0, 2},
			output: 2,
		},
		{
			nums:   []int{-4, -3, -2},
			output: 12,
		},
		{
			nums:   []int{-1, -2, -9, -6},
			output: 108,
		},
	}

	for _, tc := range tcases {
		t.Run("maxProductNaive", func(t *testing.T) {
			if output := maxProductNaive(tc.nums); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})
		t.Run("maxProductDP", func(t *testing.T) {
			if output := maxProductDP(tc.nums); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})
	}
}
