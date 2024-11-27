package maximum_profit_in_job_scheduling

import "testing"

func Test_jobScheduling(t *testing.T) {
	tcases := []struct {
		startTime, endTime, profit []int
		output                     int
	}{
		{
			startTime: []int{1, 2, 3, 3},
			endTime:   []int{3, 4, 5, 6},
			profit:    []int{50, 10, 40, 70},
			output:    120,
		},
		{
			startTime: []int{1, 2, 3, 4, 6},
			endTime:   []int{3, 5, 10, 6, 9},
			profit:    []int{20, 20, 100, 70, 60},
			output:    150,
		},
		{
			startTime: []int{1, 1, 1},
			endTime:   []int{2, 3, 4},
			profit:    []int{5, 6, 4},
			output:    6,
		},
	}

	for _, tc := range tcases {
		if output := jobScheduling(tc.startTime, tc.endTime, tc.profit); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
