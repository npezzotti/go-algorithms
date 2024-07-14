package selection_sort

import (
	"slices"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	want := []int{1, 2, 3, 4, 5}

	SelectionSort(nums)

	if !slices.IsSorted(nums) {
		t.Fatalf("got %v, want %v", nums, want)
	}
}
