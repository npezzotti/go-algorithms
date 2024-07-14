package balanced_binary_tree

import "testing"

func TestIsBalanced(t *testing.T) {
	res := isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Right: nil,
					Left:  nil,
				},
				Right: &TreeNode{
					Val:   4,
					Right: nil,
					Left:  nil,
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	})

	if res != false {
		t.Fatalf("got %v, want %v\n", res, true)
	}
}
