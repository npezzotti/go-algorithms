package build_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	mid := indexOf(inorder, root.Val)
	if mid == -1 {
		return nil
	}

	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

func indexOf(nums []int, num int) int {
	for idx, n := range nums {
		if n == num {
			return idx
		}
	}

	return -1
}
