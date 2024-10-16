package can_partition

import "testing"

func TestCanPartition(t *testing.T) {
	tcases := []struct {
		input  []int
		result bool
	}{
		{
			input:  []int{1, 5, 11, 5},
			result: true,
		},
		{
			input:  []int{1, 2, 5},
			result: false,
		},
		{
			input:  []int{1, 2, 3, 5},
			result: false,
		},
	}

	for _, tc := range tcases {
		if res := canPartition(tc.input); res != tc.result {
			t.Fatalf("expected %v for %v, got %v", tc.result, tc.input, res)
		}
	}
}
