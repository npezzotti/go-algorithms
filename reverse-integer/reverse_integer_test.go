package reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
	tcases := []struct {
		input, output int
	}{
		{
			input:  123,
			output: 321,
		},
		{
			input:  -123,
			output: -321,
		},
		{
			input:  120,
			output: 21,
		},
		{
			input:  1534236469,
			output: 0,
		},
	}

	for _, tt := range tcases {
		if output := reverse(tt.input); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
