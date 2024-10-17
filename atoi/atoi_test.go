package atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tcases := []struct {
		input  string
		output int
	}{
		// {
		// 	input:  " ",
		// 	output: 0,
		// },
		// {
		// 	input:  "42",
		// 	output: 42,
		// },
		// {
		// 	input:  "   -042",
		// 	output: -42,
		// },
		// {
		// 	input:  "1337c0d3",
		// 	output: 1337,
		// },
		// {
		// 	input:  "0-1",
		// 	output: 0,
		// },
		// {
		// 	input:  "words and 987",
		// 	output: 0,
		// },
		// {
		// 	input:  "-91283472332",
		// 	output: -2147483648,
		// },
		{
			input:  "9223372036854775808",
			output: 2147483647,
		},
		// {
		// 	input:  "+-12",
		// 	output: 0,
		// },
	}

	for _, tc := range tcases {
		if output := myAtoi(tc.input); output != tc.output {
			t.Fatalf("got %d, expected %d", output, tc.output)
		}
	}
}
