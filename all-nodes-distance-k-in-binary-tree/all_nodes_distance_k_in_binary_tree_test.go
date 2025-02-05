package all_nodes_distance_k_in_binary_tree

import (
	"slices"
	"testing"
)

func Test_distanceK(t *testing.T) {
	tcases := []struct {
		root      *TreeNode
		target, k int
		output    []int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			target: 5,
			k:      2,
			output: []int{7, 4, 1},
		},
		{
			root:   &TreeNode{Val: 1},
			target: 1,
			k:      3,
			output: []int{},
		},
	}

	for _, tt := range tcases {
		if output := distanceK(tt.root, findNodeByValue(tt.root, tt.target), tt.k); !slices.Equal(output, tt.output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}

// helper function to find the given target in the tree
func findNodeByValue(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == value {
		return root
	}

	left := findNodeByValue(root.Left, value)
	if left != nil {
		return left
	}

	return findNodeByValue(root.Right, value)
}
