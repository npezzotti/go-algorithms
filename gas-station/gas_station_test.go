package gas_station

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	tcases := []struct {
		gas    []int
		cost   []int
		output int
	}{
		{
			gas:    []int{1, 2, 3, 4, 5},
			cost:   []int{3, 4, 5, 1, 2},
			output: 3,
		},
		{
			gas:    []int{2, 3, 4},
			cost:   []int{3, 4, 3},
			output: -1,
		},
	}

	for _, tc := range tcases {
		if output := canCompleteCircuit(tc.gas, tc.cost); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
