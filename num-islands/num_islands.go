package num_islands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	var islands int

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == '1' {
			grid[row][col] = '0'
			dfs(row-1, col)
			dfs(row+1, col)
			dfs(row, col-1)
			dfs(row, col+1)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				dfs(r, c)
				islands++
			}
		}
	}

	return islands
}
