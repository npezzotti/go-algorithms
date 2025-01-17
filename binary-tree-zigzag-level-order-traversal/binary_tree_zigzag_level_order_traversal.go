package binary_tree_zigzag_level_order_traversal

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}

	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	var startLeft bool
	for len(queue) > 0 {
		set := make([]int, 0)
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			set = append(set, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if startLeft {
			slices.Reverse(set)
		}

		startLeft = !startLeft
		res = append(res, set)
	}

	return res
}
