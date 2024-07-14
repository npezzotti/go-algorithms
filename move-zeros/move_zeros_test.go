package move_zeros

import "testing"

func TestMoveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	want := []int{1, 3, 12, 0, 0}
	if len(nums) != len(want) {
		t.Fatalf("got %v, want %v", nums, want)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != want[i] {
			t.Fatalf("got %v, want %v", nums, want)
		}
	}
}
