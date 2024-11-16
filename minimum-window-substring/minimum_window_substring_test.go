package minimum_window_substring

import "testing"

func Test_min(t *testing.T) {
	tcases := []struct {
		name   string
		s      string
		t      string
		output string
	}{
		{
			name:   "regular string",
			s:      "ADOBECODEBANC",
			t:      "ABC",
			output: "BANC",
		},
		{
			name:   "single letter",
			s:      "a",
			t:      "a",
			output: "a",
		},
		{
			name:   "no window",
			s:      "a",
			t:      "aa",
			output: "",
		},
		{
			name:   "duplicate characters in string",
			s:      "bba",
			t:      "ab",
			output: "ba",
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			if output := minWindow(tc.s, tc.t); output != tc.output {
				t.Errorf("got %q, expected %q", output, tc.output)
			}
		})
	}
}
