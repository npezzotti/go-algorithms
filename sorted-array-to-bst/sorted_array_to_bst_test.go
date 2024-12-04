package sorted_array_to_bst

import (
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	tcases := []struct {
		nums   []int
		output *TreeNode
	}{
		{
			nums: []int{-10, -3, 0, 5, 9},
			output: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		{
			nums: []int{1, 3},
			output: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
	}

	for _, tc := range tcases {
		if node := sortedArrayToBST(tc.nums); !IsEqual(node, tc.output) {
			t.Errorf("unexpected tree for %v", tc.nums)
		}
	}
}

func IsEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val &&
		IsEqual(a.Left, b.Left) &&
		IsEqual(a.Right, b.Right)
}
