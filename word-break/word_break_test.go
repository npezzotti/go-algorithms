package word_break

import "testing"

func TestWordBreak(t *testing.T) {
	tcases := map[string]struct {
		inputString string
		wordDict    []string
		output      bool
	}{
		"\"leetcode\" can be segmented as \"leet code\"": {
			inputString: "leetcode",
			wordDict:    []string{"leet", "code"},
			output:      true,
		},
		"\"applepenapple\" can be segmented as \"apple pen apple\".": {
			inputString: "applepenapple",
			wordDict:    []string{"apple","pen"},
			output:      true,
		},
		"\"catsandog\" can't be segmented into available words": {
			inputString: "catsandog",
			wordDict:    []string{"cats","dog","sand","and","cat"},
			output:      false,
		},
	}

	for tname, tc := range tcases {
		t.Run(tname, func(t *testing.T) {
			if output := wordBreak(tc.inputString, tc.wordDict); output != tc.output {
				t.Fatalf("expected %v, got %v", tc.output, output)
			}
		})
	}
}
