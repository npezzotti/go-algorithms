package right_side_view

import (
	"slices"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tcases := []struct {
		tree   *TreeNode
		output []int
	}{
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			output: []int{1, 3, 4},
		},
		{
			tree: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			output: []int{1, 3},
		},
		{
			tree: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			output: []int{1, 2},
		},
		{
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			output: []int{1, 3, 4},
		},
	}

	t.Run("rightSideViewDFS", func(t *testing.T) {
		for _, tc := range tcases {
			if output := rightSideViewDFS(tc.tree); !slices.Equal(tc.output, output) {
				t.Errorf("rightSideViewDFS: expected %+v, got %+v", tc.output, output)
			}
		}
	})

	t.Run("rightSideViewBFS", func(t *testing.T) {
		for _, tc := range tcases {
			if output := rightSideViewBFS(tc.tree); !slices.Equal(tc.output, output) {
				t.Errorf("rightSideViewBFS: expected %+v, got %+v", tc.output, output)
			}
		}
	})
}
