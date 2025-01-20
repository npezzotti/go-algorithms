package path_sum_III

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	prefixSums := make(map[int]int)
	prefixSums[0]++

	var dfs func(node *TreeNode, curSum int) int
	dfs = func(node *TreeNode, curSum int) int {
		if node == nil {
			return 0
		}

		curSum += node.Val

		var pathCount int
		if count, ok := prefixSums[curSum-targetSum]; ok {
			pathCount += count
		}

		prefixSums[curSum]++

		pathCount += dfs(node.Left, curSum)
		pathCount += dfs(node.Right, curSum)

		prefixSums[curSum]--

		return pathCount
	}

	return dfs(root, 0)
}
