package word_break

import (
	"slices"
)

func wordBreak(s string, wordDict []string) bool {
	wordTracker := make([]bool, len(s)+1)
	wordTracker[0] = true

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if wordTracker[j] && slices.Contains(wordDict, string(s[j:i])) {
				wordTracker[i] = true
				break
			}
		}
	}
	return wordTracker[len(s)]
}
