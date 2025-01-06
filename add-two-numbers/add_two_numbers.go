package add_two_numbers

import (
	"github.com/npezzotti/go-algorithms/utils"
)

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	res := &utils.ListNode{}
	curNode := res

	var carry bool
	for l1 != nil || l2 != nil || carry {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
		}

		if l2 != nil {
			val2 = l2.Val
		}

		sum := val1 + val2

		if carry {
			sum += 1
			carry = false
		}

		if sum > 9 {
			carry = true
			sum = sum - 10
		}

		curNode.Next = &utils.ListNode{
			Val: sum,
		}

		curNode = curNode.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return res.Next
}
