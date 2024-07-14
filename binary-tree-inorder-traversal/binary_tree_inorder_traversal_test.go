package binary_tree_inorder_traversal

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left:  nil,
				Right: nil,
				Val:   1,
			},
			Right: &TreeNode{
				Left: nil,
				Right: &TreeNode{
					Left:  nil,
					Right: nil,
					Val:   4,
				},
				Val: 2,
			},
			Val: 5,
		},
		Right: &TreeNode{
			Left: nil,
			Right: &TreeNode{
				Left: &TreeNode{
					Left: nil,
					Right: &TreeNode{
						Left:  nil,
						Right: nil,
						Val:   5,
					},
					Val: 7,
				},
				Right: nil,
				Val:   9,
			},
			Val: 7,
		},
		Val: 1,
	}
	got := inorderTraversal(root)
	want := []int{1, 5, 2, 4, 1, 7, 7, 5, 9}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
