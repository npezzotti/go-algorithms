package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricRecursive(root *TreeNode) bool {
	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}

func isSymmetricIterative(root *TreeNode) bool {
	var lQueue []*TreeNode
	var rQueue []*TreeNode
	lQueue = append(lQueue, root.Left)
	rQueue = append(rQueue, root.Right)

	for len(lQueue) > 0 && len(rQueue) > 0 {
		lNode := lQueue[len(lQueue)-1]
		lQueue = lQueue[:len(lQueue)-1]
		rNode := rQueue[len(rQueue)-1]
		rQueue = rQueue[:len(rQueue)-1]

		if lNode == nil && rNode == nil {
			continue
		}

		if lNode == nil || rNode == nil || lNode.Val != rNode.Val {
			return false
		}

		lQueue = append(lQueue, lNode.Left, lNode.Right)
		rQueue = append(rQueue, rNode.Right, rNode.Left)
	}

	return true
}
