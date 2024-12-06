package daily_temperatures

import (
	"slices"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tcases := []struct {
		temperatures []int
		output       []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			output:       []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			output:       []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			output:       []int{1, 1, 0},
		},
	}

	for _, tc := range tcases {
		if output := dailyTemperatures(tc.temperatures); !slices.Equal(output, tc.output) {
			t.Errorf("got %v for %v, expected %v", output, tc.temperatures, tc.output)
		}
	}
}
