package zero_one_matrix

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return [][]int{}
	}

	h, w := len(mat), len(mat[0])
	queue := make([][]int, 0)
	max := h * w

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = max
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			row, col := cell[0]+direction[0], cell[1]+direction[1]
			if row >= 0 && row < h && col >= 0 && col < w && mat[row][col] > mat[cell[0]][cell[1]] {
				queue = append(queue, []int{row, col})
				mat[row][col] = mat[cell[0]][cell[1]] + 1
			}
		}
	}
	return mat
}
