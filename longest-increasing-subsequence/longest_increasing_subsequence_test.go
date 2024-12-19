package longest_increasing_subsequence

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			output: 4,
		},
		{
			nums:   []int{0, 1, 0, 3, 2, 3},
			output: 4,
		},
		{
			nums:   []int{7, 7, 7, 7, 7, 7, 7},
			output: 1,
		},
	}

	for _, tc := range tcases {
		t.Run("lengthOfLISDP", func(t *testing.T) {
			if output := lengthOfLISDP(tc.nums); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})

		t.Run("lengthOfLISBinarySearch", func(t *testing.T) {
			if output := lengthOfLISBinarySearch(tc.nums); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})
	}
}
