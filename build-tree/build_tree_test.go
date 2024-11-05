package build_tree

import "testing"

func Test_buildTree(t *testing.T) {
	tcases := []struct {
		preorder []int
		inorder  []int
		tree     *TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			tree: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			tree:     &TreeNode{Val: -1},
		},
	}

	for _, tc := range tcases {
		if tree := buildTree(tc.preorder, tc.inorder); !Equal(tree, tc.tree) {
			t.Errorf("expected %+v, got %+v", tc.tree, tree)
		}
	}
}

func Equal(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return Equal(t1.Left, t2.Left) && Equal(t1.Right, t2.Right)
}
