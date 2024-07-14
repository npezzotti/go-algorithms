package sqrt

import "testing"

func TestMySqrt(t *testing.T) {
	got := mySqrt(9)
	want := 3

	if got != want {
		t.Fatalf("want %d, got %d", got, want)
	}
}
