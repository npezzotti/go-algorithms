package word_ladder

import "testing"

func Test_ladderLength(t *testing.T) {
	tcases := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		output    int
	}{
		{
			name:      "5 permutations",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			output:    5,
		},
		{
			name:      "end word not in list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			output:    0,
		},
		{
			name:      "single characters",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			output:    2,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			if output := ladderLength(tc.beginWord, tc.endWord, tc.wordList); output != tc.output {
				t.Errorf("got %d, expected %d", output, tc.output)
			}
		})
	}
}
