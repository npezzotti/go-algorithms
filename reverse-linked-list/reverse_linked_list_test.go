package reverse_linked_list

import "testing"

func TestReverseLinkedList(t *testing.T) {
	got := reverseList(&ListNode{
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
	})

	if got.Val != 5 {
		t.Fatal("")
	}
}
