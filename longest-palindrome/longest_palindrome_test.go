package longest_palindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		name, s string
		want    int
	}{
		{"not all chars make pair", "Aabccba", 7},
		{"longest palindrome equals len(s)", "abcbca", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			if got != tt.want {
				t.Fatalf("got %d, want %d\n", got, tt.want)
			}
		})
	}
}
