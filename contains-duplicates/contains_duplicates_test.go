package contains_duplicates

import "testing"

func TestContainsDuplicate(t *testing.T) {
	got := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
