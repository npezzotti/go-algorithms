package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev, cur, next *ListNode
	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
