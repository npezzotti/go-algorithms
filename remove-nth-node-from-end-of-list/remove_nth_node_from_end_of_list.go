package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l, r := head, head

	for n > 0 {
		r = r.Next
		n--
	}

	if r == nil {
		return head.Next
	}

	for r.Next != nil {
		l, r = l.Next, r.Next
	}
	
	l.Next = l.Next.Next

	return head
}
