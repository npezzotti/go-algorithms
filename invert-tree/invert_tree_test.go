package invert_tree

import (
	"testing"
)

func TestInvertTree(t *testing.T) {
	got := InvertTree(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	})

	want := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	if !TreesEqual(got, want) {
		t.Fatalf("got %+v, want %+v\n", got, want)
	}
}

func TreesEqual(a *TreeNode, b *TreeNode) bool {
	if a != nil && b != nil {
		return a.Val == b.Val && TreesEqual(a.Right, b.Right) && TreesEqual(a.Left, b.Left)
	} else if a == nil && b == nil {
		return true
	} else {
		return false
	}
}
