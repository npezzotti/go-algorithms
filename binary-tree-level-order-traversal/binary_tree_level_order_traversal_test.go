package binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {
	tcases := map[string]struct {
		root *TreeNode
		want [][]int
	}{
		"multiple levels": {
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		"one level": {
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		"nil root": {
			root: nil,
			want: [][]int{},
		},
	}

	for tn, tc := range tcases {
		t.Run(tn, func(t *testing.T) {
			got := levelOrder(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}
