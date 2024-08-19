package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, 0)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				level = append(level, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}

		if len(level) > 0 {
			res = append(res, level)
		}
	}

	return res
}
