package word_search

import "testing"

func Test_exist(t *testing.T) {
	tcases := []struct {
		board  [][]byte
		word   string
		output bool
	}{
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "ABCCED",
			output: true,
		},
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "SEE",
			output: true,
		},
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "ABCB",
			output: false,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.word, func(t *testing.T) {
			if output := exist(tc.board, tc.word); output != tc.output {
				t.Errorf("expected %v, got %v", tc.output, output)
			}
		})
	}
}
