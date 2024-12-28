package path_sum_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	var dfs func(node *TreeNode, curSum int, curPath []int)
	dfs = func(node *TreeNode, curSum int, curPath []int) {
		if node != nil {
			curPath = append(curPath, node.Val)

			if node.Left == nil && node.Right == nil && node.Val == curSum {
				temp := make([]int, len(curPath))
				copy(temp, curPath)
				res = append(res, temp)
				return
			}

			dfs(node.Left, curSum-node.Val, curPath)
			dfs(node.Right, curSum-node.Val, curPath)
		}
	}

	if root != nil {
		dfs(root, targetSum, []int{})
	}

	return res
}
