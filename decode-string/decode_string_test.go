package decode_string

import "testing"

func Test_decodeString(t *testing.T) {
	tcases := []struct {
		s, output string
	}{
		{
			s:      "3[a]2[bc]",
			output: "aaabcbc",
		},
		{
			s:      "3[a2[c]]",
			output: "accaccacc",
		},
		{
			s:      "2[abc]3[cd]ef",
			output: "abcabccdcdcdef",
		},
		{
			s:      "abc3[cd]xyz",
			output: "abccdcdcdxyz",
		},
		{
			s:      "100[leetcode]",
			output: "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		},
	}

	for _, tt := range tcases {
		if output := decodeString(tt.s); output != tt.output {
			t.Errorf("got %s, expected %s", output, tt.output)
		}
	}
}
