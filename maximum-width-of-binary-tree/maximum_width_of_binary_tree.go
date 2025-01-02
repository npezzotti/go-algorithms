package maximum_width_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueItem struct {
	node *TreeNode
	pos  int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxWidth int

	queue := []queueItem{{root, 1}}
	for len(queue) > 0 {
		n := len(queue)
		maxWidth = max(maxWidth, queue[n-1].pos-queue[0].pos+1)
		for i := 0; i < n; i++ {
			node, pos := queue[0].node, queue[0].pos
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, queueItem{
					node: node.Left,
					pos:  2 * pos,
				})
			}

			if node.Right != nil {
				queue = append(queue, queueItem{
					node: node.Right,
					pos:  2*pos + 1,
				})
			}
		}
	}

	return maxWidth
}
