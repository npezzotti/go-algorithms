package binary_search

import "testing"

func TestSearch(t *testing.T) {
	got := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	want := 4

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
