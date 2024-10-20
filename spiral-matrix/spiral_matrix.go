package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	top, bottom, left, right := 0, rows-1, 0, cols-1
	res := make([]int, 0)

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		right--

		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}

		bottom--

		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}

		left++

		if left > right {
			break
		}
	}

	return res
}
