package rotate_square

func rotate(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])

	top, bottom := 0, rows-1
	for top < bottom {
		for i := range cols {
			matrix[top][i], matrix[bottom][i] = matrix[bottom][i], matrix[top][i]
		}

		top++
		bottom--
	}

	for i := range matrix {
		for j := i; j < cols; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
