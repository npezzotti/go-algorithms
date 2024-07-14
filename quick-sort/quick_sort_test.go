package quick_sort

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{5, 1, 3, 8, 2, 4}
	want := []int{1, 2, 3, 4, 5, 8}

	QuickSort(nums, 0, len(nums)-1)

	if !slices.IsSorted(nums) {
		t.Fatalf("got %v, want %v", nums, want)
	}
}
