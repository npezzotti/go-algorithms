package maximum_sub_array

import "testing"

func TestMaxSubArray(t *testing.T) {
	got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	want := 6

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
