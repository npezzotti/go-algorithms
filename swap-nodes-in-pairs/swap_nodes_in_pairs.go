package swap_nodes_in_pairs

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func swapPairs(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	head = next

	return head
}
