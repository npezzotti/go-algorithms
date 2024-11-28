package merge_k_sorted_lists

import "github.com/npezzotti/go-algorithms/utils"

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	if len(lists) < 1 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	l := mergeKLists(lists[:mid])
	r := mergeKLists(lists[mid:])

	return merge(l, r)
}

func merge(list1, list2 *utils.ListNode) *utils.ListNode {
	d := new(utils.ListNode)
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
