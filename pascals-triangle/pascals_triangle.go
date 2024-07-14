package pascals_triangle

func generate(numRows int) [][]int {
	triangle := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		newRow := make([]int, 0)

		for j := 0; j <= i; j++ {
			value := 1

			if i > 1 && j > 0 && j < i {
				value = triangle[i-1][j] + triangle[i-1][j-1]
			}
			newRow = append(newRow, value)
		}

		triangle = append(triangle, [][]int{newRow}...)
	}

	return triangle
}
