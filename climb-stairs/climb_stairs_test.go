package climb_stairs

import "testing"

// func TestClimbStairs(t *testing.T) {
// 	got := climbStairs(3)
// 	want := 3

// 	if got != want {
// 		t.Fatalf("got %d, want %d", got, want)
// 	}
// }

func TestClimbStairsRecursive(t *testing.T) {
	got := climbStairsRecursive(4)
	want := 5

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
