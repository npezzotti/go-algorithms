package top_k_frequent_words

import (
	"slices"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tcases := []struct {
		words  []string
		k      int
		output []string
	}{
		{
			words:  []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:      2,
			output: []string{"i", "love"},
		},
		{
			words:  []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:      4,
			output: []string{"the", "is", "sunny", "day"},
		},
	}

	for _, tc := range tcases {
		if output := topKFrequent(tc.words, tc.k); !slices.Equal(output, tc.output) {
			t.Errorf("got %v, expected %v", output, tc.output)
		}
	}
}
