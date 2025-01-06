package add_two_numbers

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	res := &utils.ListNode{}
	curNode := res

	var carry int
	for l1 != nil || l2 != nil || carry > 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10
		sum %= 10

		curNode.Next = &utils.ListNode{Val: sum}
		curNode = curNode.Next
	}

	return res.Next
}
