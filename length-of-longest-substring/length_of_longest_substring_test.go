package length_of_longest_substring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		str  string
		want int
	}{
		{
			str:  "abcabcbb",
			want: 3,
		},
		{
			str:  "bbbbb",
			want: 1,
		},
		{
			str:  "pwwkew",
			want: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tc.str); got != tc.want {
				t.Fatalf("got %d, want %d\n", got, tc.want)
			}
		})
	}
}
