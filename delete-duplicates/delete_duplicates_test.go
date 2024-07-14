package delete_duplicates

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	listHead := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}}

	got := listToArray(deleteDuplicates(listHead))
	want := listToArray(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}})

	if len(got) != len(want) {
		t.Fatalf("want %v, got %v", want, got)
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}

func listToArray(head *ListNode) []int {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}
