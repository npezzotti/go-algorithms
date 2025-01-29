package reorder_list

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func reorderList(head *utils.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast, sep := head, head, head

	for fast != nil && fast.Next != nil {
		sep = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	sep.Next = nil

	var prev, cur, next *utils.ListNode
	cur = slow

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	list1 := head
	list2 := prev

	var flip bool
	var d = new(utils.ListNode)
	for list1 != nil && list2 != nil {
		if flip {
			d.Next = list2
			list2 = list2.Next
			flip = false
		} else {
			d.Next = list1
			list1 = list1.Next
			flip = true
		}

		d = d.Next
	}

	if list1 != nil {
		d.Next = list1
	} else {
		d.Next = list2
	}

	d = nil
}
