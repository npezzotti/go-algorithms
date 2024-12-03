package symmetric_tree

import "testing"

func Test_isSymmetric(t *testing.T) {
	tcases := []struct {
		tree *TreeNode
		res  bool
	}{
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			res: true,
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			res: false,
		},
	}

	for _, tc := range tcases {
		t.Run("isSymmetricRecursive", func(t *testing.T) {
			if res := isSymmetricRecursive(tc.tree); res != tc.res {
				t.Errorf("got %t, expected %t", res, tc.res)
			}
		})

		t.Run("isSymmetricIterative", func(t *testing.T) {
			if res := isSymmetricIterative(tc.tree); res != tc.res {
				t.Errorf("got %t, expected %t", res, tc.res)
			}
		})
	}
}
