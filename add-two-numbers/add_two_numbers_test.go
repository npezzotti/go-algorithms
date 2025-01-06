package add_two_numbers

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_addTwoNumbers(t *testing.T) {
	tcases := []struct {
		l1     *utils.ListNode
		l2     *utils.ListNode
		output *utils.ListNode
	}{
		{
			l1: &utils.ListNode{
				Val: 2,
				Next: &utils.ListNode{
					Val:  4,
					Next: &utils.ListNode{Val: 3},
				},
			},
			l2: &utils.ListNode{
				Val: 5,
				Next: &utils.ListNode{
					Val:  6,
					Next: &utils.ListNode{Val: 4},
				},
			},
			output: &utils.ListNode{
				Val: 7,
				Next: &utils.ListNode{
					Val:  0,
					Next: &utils.ListNode{Val: 8},
				},
			},
		},
		{
			l1:     &utils.ListNode{Val: 0},
			l2:     &utils.ListNode{Val: 0},
			output: &utils.ListNode{Val: 0},
		},
		{
			l1: &utils.ListNode{
				Val: 9,
				Next: &utils.ListNode{
					Val: 9,
					Next: &utils.ListNode{
						Val: 9,
						Next: &utils.ListNode{
							Val: 9,
							Next: &utils.ListNode{
								Val: 9,
								Next: &utils.ListNode{
									Val: 9,
									Next: &utils.ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			l2: &utils.ListNode{
				Val: 9,
				Next: &utils.ListNode{
					Val: 9,
					Next: &utils.ListNode{
						Val: 9,
						Next: &utils.ListNode{
							Val: 9,
						},
					},
				},
			},
			output: &utils.ListNode{
				Val: 8,
				Next: &utils.ListNode{
					Val: 9,
					Next: &utils.ListNode{
						Val: 9,
						Next: &utils.ListNode{
							Val: 9,
							Next: &utils.ListNode{
								Val: 0,
								Next: &utils.ListNode{
									Val: 0,
									Next: &utils.ListNode{
										Val: 0,
										Next: &utils.ListNode{
											Val: 1,
										},
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
		if output := addTwoNumbers(tt.l1, tt.l2); !utils.IsEqual(output, tt.output) {
			t.Errorf("nodes are not equal")
		}
	}
}
