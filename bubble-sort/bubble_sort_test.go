package bubble_sort

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{7, 6, 3, 2, 5}
	want := []int{2, 3, 5, 6, 7}

	BubbleSort(nums)

	if !slices.IsSorted(nums) {
		t.Fatalf("got %v, want %v", nums, want)
	}
}
