package same_tree

import "testing"

func TestIsSameTree(t *testing.T) {
	p := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}

	q := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}

	if !isSameTree(p, q) {
		t.Errorf("unexpected failure")
	}
}
