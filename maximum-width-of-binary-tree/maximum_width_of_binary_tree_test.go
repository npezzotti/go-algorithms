package maximum_width_of_binary_tree

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	tcases := []struct {
		root   *TreeNode
		output int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 9},
				},
			},
			output: 4,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  5,
						Left: &TreeNode{Val: 6},
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:  9,
						Left: &TreeNode{Val: 7},
					},
				},
			},
			output: 7,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: 2,
		},
	}

	for _, tt := range tcases {
		if output := widthOfBinaryTree(tt.root); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
