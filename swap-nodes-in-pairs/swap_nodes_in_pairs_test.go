package swap_nodes_in_pairs

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_swapPairs(t *testing.T) {
	tcases := []struct {
		head   *utils.ListNode
		output *utils.ListNode
	}{
		{
			head: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val: 2,
					Next: &utils.ListNode{
						Val:  3,
						Next: &utils.ListNode{Val: 4},
					},
				},
			},
			output: &utils.ListNode{
				Val: 2,
				Next: &utils.ListNode{
					Val: 1,
					Next: &utils.ListNode{
						Val:  4,
						Next: &utils.ListNode{Val: 3},
					},
				},
			},
		},
		{
			head: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val:  2,
					Next: &utils.ListNode{Val: 3},
				},
			},
			output: &utils.ListNode{
				Val: 2,
				Next: &utils.ListNode{
					Val:  1,
					Next: &utils.ListNode{Val: 3},
				},
			},
		},
		{
			head:   nil,
			output: nil,
		},
		{
			head:   &utils.ListNode{Val: 1},
			output: &utils.ListNode{Val: 1},
		},
	}

	for _, tc := range tcases {
		if output := swapPairs(tc.head); !utils.IsEqual(output, tc.output) {
			t.Error("lists are not equal")
		}
	}
}
