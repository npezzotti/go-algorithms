package path_sum

import "testing"

func Test_hasPathSum(t *testing.T) {
	tcases := []struct {
		root      *TreeNode
		targetSum int
		output    bool
	}{
		{
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   11,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			output:    true,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 5,
			output:    false,
		},
		{
			root:      nil,
			targetSum: 0,
			output:    false,
		},
		{
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			targetSum: 1,
			output:    false,
		},
	}

	for _, tt := range tcases {
		if output := hasPathSum(tt.root, tt.targetSum); output != tt.output {
			t.Errorf("got %t for %+v with sum %d, expected %t", output, tt.root, tt.targetSum, tt.output)
		}
	}
}
