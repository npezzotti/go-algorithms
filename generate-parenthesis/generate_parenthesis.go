package generate_parenthesis

func generateParenthesis(n int) []string {
	var res []string

	var dfs func(l, r int, curRes string)
	dfs = func(l, r int, curRes string) {
		if l >= n && r >= n {
			res = append(res, curRes)
			return
		}

		if l < n {
			dfs(l+1, r, curRes+"(")
		}

		if r < l {
			dfs(l, r+1, curRes+")")
		}
	}

	dfs(0, 0, "")

	return res
}
