package single_number

import "testing"

func Test_singleNumber(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{2, 2, 1},
			output: 1,
		},
		{
			nums:   []int{4, 1, 2, 1, 2},
			output: 4,
		},
		{
			nums:   []int{1},
			output: 1,
		},
	}

	for _, tc := range tcases {
		if output := singleNumber(tc.nums); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
