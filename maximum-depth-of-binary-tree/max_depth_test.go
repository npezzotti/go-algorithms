package max_depth

import "testing"

var root *TreeNode = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}

func TestMaxDepth(t *testing.T) {
	got := maxDepth(root)

	if got != 3 {
		t.Fatalf("want %d, got %d\n", 3, got)
	}
}

func TestMaxDepthIterative(t *testing.T) {
	got := maxDepthIterative(root)

	if got != 3 {
		t.Fatalf("want %d, got %d\n", 3, got)
	}
}
