package binary_tree_zigzag_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tcases := []struct {
		root   *TreeNode
		output [][]int
	}{
		{
			root: &TreeNode{
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
			output: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			root:   &TreeNode{Val: 1},
			output: [][]int{{1}},
		},
		{
			root:   nil,
			output: [][]int{},
		},
	}

	for _, tt := range tcases {
		if output := zigzagLevelOrder(tt.root); !reflect.DeepEqual(tt.output, output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
