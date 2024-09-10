package valid_binary_search_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return valid(node.Left, min, node.Val) && valid(node.Right, node.Val, max)
}
