package sort_list

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_sortList(t *testing.T) {
	tcases := []struct {
		head   *utils.ListNode
		output *utils.ListNode
	}{
		{
			head: &utils.ListNode{
				Val: 4,
				Next: &utils.ListNode{
					Val: 2,
					Next: &utils.ListNode{
						Val:  1,
						Next: &utils.ListNode{Val: 3},
					},
				},
			},
			output: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val: 2,
					Next: &utils.ListNode{
						Val:  3,
						Next: &utils.ListNode{Val: 4},
					},
				},
			},
		},
		{
			head: &utils.ListNode{
				Val: -1,
				Next: &utils.ListNode{
					Val: 5,
					Next: &utils.ListNode{
						Val: 3,
						Next: &utils.ListNode{
							Val:  4,
							Next: &utils.ListNode{Val: 0},
						},
					},
				},
			},
			output: &utils.ListNode{
				Val: -1,
				Next: &utils.ListNode{
					Val: 0,
					Next: &utils.ListNode{
						Val: 3,
						Next: &utils.ListNode{
							Val:  4,
							Next: &utils.ListNode{Val: 5},
						},
					},
				},
			},
		},
		{
			head:   nil,
			output: nil,
		},
	}

	for _, tt := range tcases {
		if output := sortList(tt.head); !utils.IsEqual(output, tt.output) {
			t.Errorf("linked lists do not match")
		}
	}
}
