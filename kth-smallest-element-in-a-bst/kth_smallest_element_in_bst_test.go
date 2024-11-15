package kth_smallest_element_in_bst

import "testing"

func Test_kthSmallest(t *testing.T) {
	tcases := []struct {
		root   *TreeNode
		k      int
		output int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 4},
			},
			k:      1,
			output: 1,
		},
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:      3,
			output: 3,
		},
	}

	for _, tc := range tcases {
		if output := kthSmallest(tc.root, tc.k); output != tc.output {
			t.Errorf("expected %d, got %d", tc.output, output)
		}
	}
}
