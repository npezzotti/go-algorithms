package maximal_square

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	var mx int
	for i, row := range matrix {
		for j, col := range row {
			if col == '1' {
				mn := min(
					dp[i+1][j],
					dp[i][j+1],
					dp[i][j],
				)

				dp[i+1][j+1] = mn + 1

				mx = max(mx, dp[i+1][j+1])

			}
		}
	}

	return mx * mx
}
