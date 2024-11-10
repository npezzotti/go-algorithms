package find_anagrams

import (
	"fmt"
	"slices"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tcases := []struct {
		s      string
		p      string
		output []int
	}{
		{
			s:      "cbaebabacd",
			p:      "abc",
			output: []int{0, 6},
		},
		{
			s:      "abab",
			p:      "ab",
			output: []int{0, 1, 2},
		},
	}

	for _, tc := range tcases {
		t.Run(fmt.Sprintf("%s in %s", tc.p, tc.s), func(t *testing.T) {
			if output := findAnagrams(tc.s, tc.p); !slices.Equal(output, tc.output) {
				t.Errorf("expected %+v, got %+v", tc.output, output)
			}
		})
	}
}
