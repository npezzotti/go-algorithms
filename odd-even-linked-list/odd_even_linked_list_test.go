package odd_even_linked_list

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_oddEvenList(t *testing.T) {
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
						Val: 3,
						Next: &utils.ListNode{
							Val: 4,
							Next: &utils.ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			output: &utils.ListNode{
				Val: 1,
				Next: &utils.ListNode{
					Val: 3,
					Next: &utils.ListNode{
						Val: 5,
						Next: &utils.ListNode{
							Val: 2,
							Next: &utils.ListNode{
								Val: 4,
							},
						},
					},
				},
			},
		},
		{
			head: &utils.ListNode{
				Val: 2,
				Next: &utils.ListNode{
					Val: 1,
					Next: &utils.ListNode{
						Val: 3,
						Next: &utils.ListNode{
							Val: 5,
							Next: &utils.ListNode{
								Val: 6,
								Next: &utils.ListNode{
									Val: 4,
									Next: &utils.ListNode{
										Val: 7,
									},
								},
							},
						},
					},
				},
			},
			output: &utils.ListNode{
				Val: 2,
				Next: &utils.ListNode{
					Val: 3,
					Next: &utils.ListNode{
						Val: 6,
						Next: &utils.ListNode{
							Val: 7,
							Next: &utils.ListNode{
								Val: 1,
								Next: &utils.ListNode{
									Val: 5,
									Next: &utils.ListNode{
										Val: 4,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tcases {
		if output := oddEvenList(tt.head); !utils.IsEqual(output, tt.output) {
			t.Errorf("lists do not match")
		}
	}
}
