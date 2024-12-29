package odd_even_linked_list

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func oddEvenList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return head
	}

	evenHead, evenPointer, oddPointer := head.Next, head.Next, head

	for evenPointer != nil && evenPointer.Next != nil {
		oddPointer.Next = oddPointer.Next.Next
		evenPointer.Next = evenPointer.Next.Next
		oddPointer = oddPointer.Next
		evenPointer = evenPointer.Next
	}

	oddPointer.Next = evenHead

	return head
}
