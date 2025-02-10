package rotate_list

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func Test_rotateRight(t *testing.T) {
	tcases := []struct {
		head   *utils.ListNode
		k      int
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
			k: 2,
			output: &utils.ListNode{
				Val: 4,
				Next: &utils.ListNode{
					Val: 5,
					Next: &utils.ListNode{
						Val: 1,
						Next: &utils.ListNode{
							Val: 2,
							Next: &utils.ListNode{
								Val: 3,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tcases {
		if output := rotateRight(tt.head, tt.k); !utils.IsEqual(output, tt.output) {
			t.Errorf("list are not equal")
		}
	}
}
