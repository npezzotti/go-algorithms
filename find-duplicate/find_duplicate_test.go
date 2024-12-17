package find_duplicate

import "testing"

func Test_funcDuplicate(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{1, 3, 4, 2, 2},
			output: 2,
		},
		{
			nums:   []int{3, 1, 3, 4, 2},
			output: 3,
		},
		{
			nums:   []int{3, 3, 3, 3, 3},
			output: 3,
		},
		{
			nums:   []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			output: 9,
		},
	}

	for _, tc := range tcases {
		if output := findDuplicate(tc.nums); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
