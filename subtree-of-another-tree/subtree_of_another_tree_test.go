package subtree_of_another_tree

import "testing"

func Test_isSubtree(t *testing.T) {
	tcases := []struct {
		root    *TreeNode
		subRoot *TreeNode
		output  bool
	}{
		// {
		// 	root: &TreeNode{
		// 		Val: 3,
		// 		Left: &TreeNode{
		// 			Val:   4,
		// 			Left:  &TreeNode{Val: 1},
		// 			Right: &TreeNode{Val: 2},
		// 		},
		// 		Right: &TreeNode{Val: 5},
		// 	},
		// 	subRoot: &TreeNode{
		// 		Val:   4,
		// 		Left:  &TreeNode{Val: 1},
		// 		Right: &TreeNode{Val: 2},
		// 	},
		// 	output: true,
		// },
		// {
		// 	root: &TreeNode{
		// 		Val: 3,
		// 		Left: &TreeNode{
		// 			Val:  4,
		// 			Left: &TreeNode{Val: 1},
		// 			Right: &TreeNode{
		// 				Val: 2,
		// 				Right: &TreeNode{
		// 					Val:  2,
		// 					Left: &TreeNode{Val: 0},
		// 				},
		// 			},
		// 		},
		// 		Right: &TreeNode{Val: 5},
		// 	},
		// 	subRoot: &TreeNode{
		// 		Val:   4,
		// 		Left:  &TreeNode{Val: 1},
		// 		Right: &TreeNode{Val: 2},
		// 	},
		// 	output: false,
		// },
		// {
		// 	root:    &TreeNode{Val: 1, Left: &TreeNode{Val: 1}},
		// 	subRoot: &TreeNode{Val: 1},
		// 	output:  true,
		// },
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 2},
				},
			},
			subRoot: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			output: false,
		},
	}

	for _, tc := range tcases {
		if output := isSubtree(tc.root, tc.subRoot); output != tc.output {
			t.Errorf("got %t, expected %t", output, tc.output)
		}
	}
}
