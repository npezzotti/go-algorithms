package longest_consecutive_sequence

import "testing"

func Test_longestConsecutive(t *testing.T) {
	tcases := []struct {
		nums   []int
		output int
	}{
		// {
		// 	nums:   []int{100, 4, 200, 1, 3, 2},
		// 	output: 4,
		// },
		{
			nums:   []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			output: 9,
		},
	}

	for _, tt := range tcases {
		t.Run("longestConsecutiveBruteForce", func(t *testing.T) {
			if output := longestConsecutiveBruteForce(tt.nums); output != tt.output {
				t.Errorf("got %d, expected %d", output, tt.output)
			}
		})

		t.Run("longestConsecutive", func(t *testing.T) {
			if output := longestConsecutive(tt.nums); output != tt.output {
				t.Errorf("got %d, expected %d", output, tt.output)
			}
		})
	}
}
