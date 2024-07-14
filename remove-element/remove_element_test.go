package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	got := RemoveElement([]int{1, 2, 3}, 1)
	want := 2

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
