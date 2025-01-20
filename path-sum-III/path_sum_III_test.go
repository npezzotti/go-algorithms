package path_sum_III

import "testing"

func Test_pathSum(t *testing.T) {
	tcases := []struct {
		root      *TreeNode
		targetSum int
		output    int
	}{
		{
			root: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: -2},
					},
					Right: &TreeNode{
						Val:   2,
						Right: &TreeNode{Val: 1},
					},
				},
				Right: &TreeNode{
					Val:   -3,
					Right: &TreeNode{Val: 11},
				},
			},
			targetSum: 8,
			output:    3,
		},
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
			output:    3,
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			},
			targetSum: 3,
			output:    2,
		},
		{
			root:      &TreeNode{Val: 1},
			targetSum: 0,
			output:    0,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: -2},
				Right: &TreeNode{Val: -3},
			},
			targetSum: -1,
			output:    1,
		},
	}

	for _, tt := range tcases {
		if output := pathSum(tt.root, tt.targetSum); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
