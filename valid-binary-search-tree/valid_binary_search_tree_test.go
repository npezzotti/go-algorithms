package valid_binary_search_tree

import "testing"

func TestIsValidBST(t *testing.T) {
	tcases := []struct {
		root  *TreeNode
		valid bool
	}{
		{
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			valid: true,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			valid: false,
		},
	}

	for _, tc := range tcases {
		if got := isValidBST(tc.root); got != tc.valid {
			t.Fatalf("expected %+v validity to be %t, got %t", tc.root, tc.valid, got)
		}
	}
}
