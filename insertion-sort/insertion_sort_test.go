package insertion_sort

import (
	"slices"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}

	InsertionSort(nums)

	if !slices.IsSorted(nums) {
		t.Fatalf("got %v, want %v", nums, want)
	}
}
