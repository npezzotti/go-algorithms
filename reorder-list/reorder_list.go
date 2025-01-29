package reorder_list

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func reorderList(head *utils.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev, cur, next *utils.ListNode
	cur = slow.Next
	slow.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	firstHalf := head
	secondHalf := prev
	for secondHalf != nil {
		temp1 := firstHalf.Next
		temp2 := secondHalf.Next
		firstHalf.Next = secondHalf
		secondHalf.Next = temp1
		firstHalf = temp1
		secondHalf = temp2
	}
}
