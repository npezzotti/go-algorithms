package linked_list_cycle

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	listNode1 := &ListNode{Val: 1, Next: nil}
	listNode2 := &ListNode{Val: 5, Next: nil}
	listNode3 := &ListNode{Val: 8, Next: nil}
	listNode1.Next = listNode2
	listNode2.Next = listNode3
	listNode3.Next = listNode1

	got := hasCycle(listNode1)
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
