package unique_paths

import "fmt"

func generateKey(m, n int) string {
	return fmt.Sprintf("%d,%d", m, n)
}

func uniquePathsRecursive(m, n int) int {
	cache := make(map[string]int)

	var recurse func(m, n int) int
	recurse = func(m, n int) int {
		if m == 1 || n == 1 {
			return 1
		}

		if v, ok := cache[generateKey(m, n)]; ok {
			return v
		}

		cache[generateKey(m, n)] = recurse(m-1, n) + recurse(m, n-1)
		return cache[generateKey(m, n)]
	}

	return recurse(m, n)
}

func uniquePathsDP(m, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
