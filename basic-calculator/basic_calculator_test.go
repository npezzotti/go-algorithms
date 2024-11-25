package basic_calculator

import "testing"

func Test_calculate(t *testing.T) {
	tcases := []struct {
		name   string
		s      string
		output int
	}{
		{
			name:   "addition",
			s:      "1 + 1",
			output: 2,
		},
		{
			name:   "addition and subtraction",
			s:      " 2-1 + 2 ",
			output: 3,
		},
		{
			name:   "ordered operations",
			s:      "(1+(4+5+2)-3)+(6+8)",
			output: 23,
		},
		{
			name:   "multi-digit numbers",
			s:      "123+56-35",
			output: 144,
		},
		{
			name:   "negative numbers",
			s:      "1-(     -2)",
			output: 3,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			if output := calculate(tc.s); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})
	}
}
