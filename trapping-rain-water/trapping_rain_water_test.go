package trapping_rain_water

import "testing"

func Test_trap(t *testing.T) {
	tcases := []struct {
		height []int
		output int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			output: 6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			output: 9,
		},
	}

	t.Run("trapBruteForce", func(t *testing.T) {
		for _, tc := range tcases {
			if output := trapBruteForce(tc.height); output != tc.output {
				t.Errorf("got %d for height %v, expected %d", output, tc.height, tc.output)
			}
		}
	})

	t.Run("trapBruteForceTwoArrs", func(t *testing.T) {
		for _, tc := range tcases {
			if output := trapBruteForceTwoArrs(tc.height); output != tc.output {
				t.Errorf("got %d for height %v, expected %d", output, tc.height, tc.output)
			}
		}
	})

	t.Run("trapStack", func(t *testing.T) {
		for _, tc := range tcases {
			if output := trapStack(tc.height); output != tc.output {
				t.Errorf("got %d for height %v, expected %d", output, tc.height, tc.output)
			}
		}
	})

	t.Run("trapTwoPointers", func(t *testing.T) {
		for _, tc := range tcases {
			if output := trapTwoPointers(tc.height); output != tc.output {
				t.Errorf("got %d for height %v, expected %d", output, tc.height, tc.output)
			}
		}
	})
}
