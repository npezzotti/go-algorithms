package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	d := new(ListNode)
	cur := d

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return d.Next
}
