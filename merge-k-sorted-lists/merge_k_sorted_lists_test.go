package merge_k_sorted_lists

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_mergekLists(t *testing.T) {
	tcases := []struct {
		lists  []*utils.ListNode
		output *utils.ListNode
	}{
		{
			lists: []*utils.ListNode{
				{
					Val: 1,
					Next: &utils.ListNode{
						Val:  4,
						Next: &utils.ListNode{Val: 5},
					},
				},
				{
					Val: 1,
					Next: &utils.ListNode{
						Val:  3,
						Next: &utils.ListNode{Val: 4},
					},
				},
				{
					Val:  2,
					Next: &utils.ListNode{Val: 6},
				},
			},
			output: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val: 1,
					Next: &utils.ListNode{
						Val: 2,
						Next: &utils.ListNode{
							Val: 3,
							Next: &utils.ListNode{
								Val: 4,
								Next: &utils.ListNode{
									Val: 4,
									Next: &utils.ListNode{
										Val: 5,
										Next: &utils.ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			lists:  []*utils.ListNode{},
			output: &utils.ListNode{},
		},
		{
			lists:  []*utils.ListNode{{}},
			output: &utils.ListNode{},
		},
	}

	for _, tc := range tcases {
		if output := mergeKLists(tc.lists); !utils.IsEqual(tc.output, output) {
			t.Errorf("lists are not equal")
		}
	}
}
