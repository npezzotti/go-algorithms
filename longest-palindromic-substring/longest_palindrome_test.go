package longest_palindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tcases := []struct {
		input  string
		output string
	}{
		{
			input:  "babad",
			output: "bab",
		},
		{
			input:  "cbbd",
			output: "bb",
		},
	}

	for _, tc := range tcases {
		if output := longestPalindrome(tc.input); output != tc.output {
			t.Errorf("expected %s, got %s", tc.output, output)
		}
	}
}

func TestExpandFromCenter(t *testing.T) {
	tcases := []struct {
		l, r int
		str  string
		res  int
	}{
		{
			l:   0,
			r:   0,
			str: "a",
			res: 1,
		},
		{
			l:   3,
			r:   3,
			str: "drdad",
			res: 3,
		},
		{
			l:   1,
			r:   2,
			str: "cbbd",
			res: 2,
		},
	}

	for _, tc := range tcases {
		if res := expandFromCenter(tc.l, tc.r, tc.str); res != tc.res {
			t.Errorf("expected %d, got %d", tc.res, res)
		}
	}
}
