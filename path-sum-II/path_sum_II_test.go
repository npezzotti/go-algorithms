package path_sum_II

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	tcases := []struct {
		root      *TreeNode
		targetSum int
		output    [][]int
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
			output:    [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}

	for _, tt := range tcases {
		if output := pathSum(tt.root, tt.targetSum); !reflect.DeepEqual(output, tt.output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
