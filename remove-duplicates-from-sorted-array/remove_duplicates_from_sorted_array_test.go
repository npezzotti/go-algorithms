package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	list := []int{1, 3, 3, 6, 6, 6, 8, 8, 9}

	got := RemoveDuplicates(list)
	want := 5

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
