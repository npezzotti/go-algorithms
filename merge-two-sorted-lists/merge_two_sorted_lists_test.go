package merge_two_sorted_lists

import (
	"testing"

	"github.com/npezzotti/go-algorithms/utils"
)

func TestMergeTwoLists(t *testing.T) {
	listOneHead := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 2, Next: &utils.ListNode{Val: 4, Next: nil}}}
	listTwoHead := &utils.ListNode{Val: 1, Next: &utils.ListNode{Val: 3, Next: &utils.ListNode{Val: 4, Next: nil}}}
	merged := MergeTwoLists(listOneHead, listTwoHead)

	got := listToArray(merged)
	want := []int{1, 1, 2, 3, 4, 4}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}

func TestMergeTwoListsBothEmpty(t *testing.T) {
	merged := MergeTwoLists(nil, nil)

	got := listToArray(merged)
	want := []int{}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}

func TestMergeTwoListsOneEmpty(t *testing.T) {
	listOneHead := &utils.ListNode{Val: 0, Next: nil}
	merged := MergeTwoLists(listOneHead, nil)

	got := listToArray(merged)
	want := []int{0, 1, 2, 4, 5, 9}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}

func listToArray(head *utils.ListNode) []int {
	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}
