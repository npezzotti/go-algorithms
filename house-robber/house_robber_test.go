package house_robber

import "testing"

func Test_rob(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{1, 2, 3, 1},
			output: 4,
		},
		{
			nums:   []int{2, 7, 9, 3, 1},
			output: 12,
		},
		{
			nums:   []int{2, 1, 1, 2},
			output: 4,
		},
		{
			nums:   []int{1, 3, 1, 3, 100},
			output: 103,
		},
		{
			nums:   []int{1, 2},
			output: 2,
		},
	}

	for _, tc := range tcases {
		if output := rob(tc.nums); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
