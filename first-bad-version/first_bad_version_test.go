package first_bad_version

import "testing"

func TestFirstBadVersion(t *testing.T) {
	got := firstBadVersion(10)

	if got != firstBadVersionNum {
		t.Fatalf("got %d, want %d\n", got, firstBadVersionNum)
	}
}
