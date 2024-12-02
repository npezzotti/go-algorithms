package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}

	return true
}

func isPalindromeOptimized(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev, cur, next *ListNode
	cur = slow
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	rightHead := prev
	for rightHead != nil {
		if head.Val != rightHead.Val {
			return false
		}

		head = head.Next
		rightHead = rightHead.Next
	}

	return true
}
