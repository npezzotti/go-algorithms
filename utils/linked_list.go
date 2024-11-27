package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}

		a = a.Next
		b = b.Next
	}

	if a != nil && b != nil {
		return false
	}

	return true
}
