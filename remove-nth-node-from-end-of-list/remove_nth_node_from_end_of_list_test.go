package remove_nth_node_from_end_of_list

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	tcases := []struct {
		head   *ListNode
		n      int
		output *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			n: 2,
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
		{
			head: &ListNode{
				Val: 1,
			},
			n:      1,
			output: nil,
		},
		{
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
			n:      1,
			output: &ListNode{Val: 1},
		},
		{
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
			n:      2,
			output: &ListNode{Val: 2},
		},
	}

	for _, tc := range tcases {
		if output := removeNthFromEnd(tc.head, tc.n); !IsEqual(output, tc.output) {
			t.Error("linked lists are not equal")
		}
	}
}

func IsEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	if a != nil || b != nil {
		return false
	}

	return true
}
