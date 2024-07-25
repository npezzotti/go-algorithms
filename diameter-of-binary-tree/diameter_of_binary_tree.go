package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int = 0

	var maxDepth func(n *TreeNode) int
	maxDepth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}

		left, right := maxDepth(n.Left), maxDepth(n.Right)
		maxDiameter = max(maxDiameter, left+right)

		return 1 + max(left, right)
	}

	maxDepth(root)

	return maxDiameter
}
