package hamming_weight

import "testing"

func Test_hammingWeight(t *testing.T) {
	tcases := []struct {
		input  int
		output int
	}{
		{
			input:  11,
			output: 3,
		},
		{
			input:  128,
			output: 1,
		},
		{
			input:  2147483645,
			output: 30,
		},
	}

	for _, tc := range tcases {
		if output := hammingWeight(tc.input); output != tc.output {
			t.Errorf("got %d, expected %d", output, tc.output)
		}
	}
}
