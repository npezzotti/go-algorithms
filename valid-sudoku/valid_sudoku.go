package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		set := make(map[byte]struct{}, 9)
		for j := 0; j < 9; j++ {
			if _, ok := set[board[i][j]]; ok {
				return false
			}

			if board[i][j] != '.' {
				set[board[i][j]] = struct{}{}
			}
		}
	}

	for i := 0; i < 9; i++ {
		set := make(map[byte]struct{}, 9)
		for j := 0; j < 9; j++ {
			if _, ok := set[board[j][i]]; ok {
				return false
			}

			if board[j][i] != '.' {
				set[board[j][i]] = struct{}{}
			}
		}
	}

	starts := [][2]int{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}

	for _, s := range starts {
		set := make(map[byte]struct{}, 9)
		for i := s[0]; i < s[0]+3; i++ {
			for j := s[1]; j < s[1]+3; j++ {
				if _, ok := set[board[i][j]]; ok {
					return false
				}

				if board[i][j] != '.' {
					set[board[i][j]] = struct{}{}
				}
			}
		}
	}

	return true
}
