package sort_list

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func sortList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *utils.ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	left := sortList(head)
	right := sortList(slow)

	return MergeTwoLists(left, right)
}

func MergeTwoLists(l1, l2 *utils.ListNode) *utils.ListNode {
	d := new(utils.ListNode)
	cur := d

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return d.Next
}
