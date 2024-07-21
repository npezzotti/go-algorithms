package max_depth

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive solution
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

// iterative solution
func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	_ = queue.PushBack(root)
	var level int = 0

	for queue.Len() > 0 {
		n := queue.Len()

		for i := 0; i < n; i++ {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			if node.Left != nil {
				_ = queue.PushBack(node.Left)
			}
			if node.Right != nil {
				_ = queue.PushBack(node.Right)
			}
		}

		level++
	}

	return level
}
