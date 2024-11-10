package word_search

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var backtrack func(r, c, i int) bool
	backtrack = func(r, c, i int) bool {
		if len(word) == i {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[i] {
			return false
		}

		tmp := board[r][c]
		board[r][c] = '#'
		if backtrack(r+1, c, i+1) || backtrack(r-1, c, i+1) || backtrack(r, c+1, i+1) || backtrack(r, c-1, i+1) {
			return true
		}
		board[r][c] = tmp

		return false
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				if backtrack(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}
