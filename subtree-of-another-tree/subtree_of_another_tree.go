package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == subRoot.Val {
		if isEqual(root.Left, subRoot.Left) &&
			isEqual(root.Right, subRoot.Right) {
			return true
		}
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}
