package container_with_most_water

import "testing"

func Test_maxArea(t *testing.T) {
	tcases := []struct {
		height []int
		output int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			height: []int{1, 1},
			output: 1,
		},
	}

	for _, tc := range tcases {
		if output := maxArea(tc.height); output != tc.output {
			t.Errorf("expected %d, got %d", tc.output, output)
		}
	}
}
