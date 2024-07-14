package merge_sorted_array

import "testing"

func TestMerge(t *testing.T) {
	nums := []int{2, 2, 3, 0, 0, 0}
	merge(nums, 3, []int{1, 5, 6}, 3)

	want := []int{1, 2, 2, 3, 5, 6}

	if len(nums) != len(want) {
		t.Fatalf("got %v, want %v", nums, want)
	}

	for i := range nums {
		if nums[i] != want[i] {
			t.Fatalf("got %v, want %v", nums, want)
		}
	}
}
