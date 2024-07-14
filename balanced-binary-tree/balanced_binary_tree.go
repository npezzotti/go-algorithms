package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := maxDepth(root.Left) - maxDepth(root.Right)

	if diff <= 1 && diff >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
