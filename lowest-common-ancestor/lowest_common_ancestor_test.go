package lowest_common_ancestor

import "testing"

func TestLowestCommonAcestor(t *testing.T) {
	treeNode := &TreeNode{
		Val: 6,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val:   5,
					Right: nil,
					Left:  nil,
				},
				Left: &TreeNode{
					Val:   3,
					Right: nil,
					Left:  nil,
				},
			},
			Left: &TreeNode{
				Val: 0,
			},
		},
		Left: &TreeNode{
			Val: 8,
			Right: &TreeNode{
				Val:   9,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
		},
	}

	got := lowestCommonAncestor(treeNode, treeNode.Left, treeNode.Right)

	if got != treeNode {
		t.Fatalf("got %+v, want %+v\n", got, treeNode)
	}
}
