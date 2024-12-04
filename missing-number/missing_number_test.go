package missing_number

import "testing"

func Test_missingNumber(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{3, 0, 1},
			output: 2,
		},
		{
			nums:   []int{0, 1},
			output: 2,
		},
		{
			nums:   []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			output: 8,
		},
	}

	for _, tc := range tcases {
		t.Run("missingNumberHash", func(t *testing.T) {
			if output := missingNumberHash(tc.nums); output != tc.output {
				t.Errorf("got %d for nums %v, expected %d", output, tc.nums, tc.output)
			}
		})

		t.Run("missingNumberSum", func(t *testing.T) {
			if output := missingNumberSum(tc.nums); output != tc.output {
				t.Errorf("got %d for nums %v, expected %d", output, tc.nums, tc.output)
			}
		})

		t.Run("missingNumberBitwise", func(t *testing.T) {
			if output := missingNumberBitwise(tc.nums); output != tc.output {
				t.Errorf("got %d for nums %v, expected %d", output, tc.nums, tc.output)
			}
		})
	}
}
