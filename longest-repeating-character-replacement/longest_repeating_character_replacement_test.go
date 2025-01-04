package longest_repeating_character_replacement

import "testing"

func Test_characterReplacement(t *testing.T) {
	tcases := []struct {
		s         string
		k, output int
	}{
		{
			s:      "ABAB",
			k:      2,
			output: 4,
		},
		{
			s:      "AABABBA",
			k:      1,
			output: 4,
		},
	}

	for _, tt := range tcases {
		if output := characterReplacement(tt.s, tt.k); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
