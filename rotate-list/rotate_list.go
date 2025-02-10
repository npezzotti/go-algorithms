package rotate_list

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var n int
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}

	k %= n

	if k == 0 {
		return head
	}

	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}
