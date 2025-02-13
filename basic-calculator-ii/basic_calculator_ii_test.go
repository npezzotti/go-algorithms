package basic_calculator_iii

import "testing"

func Test_calculate(t *testing.T) {
	tcases := []struct {
		s      string
		output int
	}{
		{
			s:      "3+2*2",
			output: 7,
		},
		{
			s:      " 3/2 ",
			output: 1,
		},
		{
			s:      " 3+5 / 2 ",
			output: 5,
		},
		{
			s:      "0-2147483647",
			output: -2147483647,
		},
	}

	for _, tt := range tcases {
		if output := calculate(tt.s); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
