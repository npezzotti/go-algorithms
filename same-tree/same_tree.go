package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}
