package decode_ways

import "testing"

func Test_numDecodings(t *testing.T) {
	tcases := []struct {
		s      string
		output int
	}{
		{
			s:      "12",
			output: 2,
		},
		{
			s:      "226",
			output: 3,
		},
		{
			s:      "06",
			output: 0,
		},
	}

	for _, tt := range tcases {
		if output := numDecodings(tt.s); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
