package reorder_list

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_reorderList(t *testing.T) {
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
				Val: 1,
				Next: &utils.ListNode{
					Val: 4,
					Next: &utils.ListNode{
						Val:  2,
						Next: &utils.ListNode{Val: 3},
					},
				},
			},
		},
		{
			head: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val: 2,
					Next: &utils.ListNode{
						Val: 3,
						Next: &utils.ListNode{
							Val:  4,
							Next: &utils.ListNode{Val: 5},
						},
					},
				},
			},
			output: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val: 5,
					Next: &utils.ListNode{
						Val: 2,
						Next: &utils.ListNode{
							Val:  4,
							Next: &utils.ListNode{Val: 3},
						},
					},
				},
			},
		},
		{
			head:   &utils.ListNode{Val: 1},
			output: &utils.ListNode{Val: 1},
		},
	}

	for _, tt := range tcases {
		reorderList(tt.head)
		if !utils.IsEqual(tt.head, tt.output) {
			t.Errorf("lists are not equal")
		}
	}
}
