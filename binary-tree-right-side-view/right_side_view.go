package right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideViewDFS(root *TreeNode) []int {
	res := make([]int, 0)

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level == len(res) {
			res = append(res, node.Val)
		}

		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}

	dfs(root, 0)

	return res
}

func rightSideViewBFS(root *TreeNode) []int {
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		queueLen := len(queue)
		res = append(res, queue[queueLen-1].Val)

		for i := 0; i < queueLen; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return res
}
