package merge_sort

import (
	"slices"
	"testing"
)

func TestMergeSort(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 8}
	got := MergeSort([]int{5, 1, 2, 6, 8, 3, 4})

	if !slices.IsSorted(got) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
