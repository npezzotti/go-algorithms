package is_isomorphic

import "testing"

func TestIsIsomorphic(t *testing.T) {
	got := isIsomorphic("paper", "title")
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
