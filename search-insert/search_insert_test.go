package search_insert

import "testing"

func TestSearchInsertExists(t *testing.T) {
	got := SearchInsert([]int{2, 5, 6, 8, 9}, 6)
	want := 2

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestSearchInsertNotExist(t *testing.T) {
	got := SearchInsert([]int{2, 5, 6, 8, 9}, 7)
	want := 3

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
