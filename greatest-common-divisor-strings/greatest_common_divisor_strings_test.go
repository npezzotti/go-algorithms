package greatest_common_divisor_strings

import "testing"

func TestGcdOfStrings(t *testing.T) {
	got := gcdOfStrings("ABABAB", "ABAB")
	want := "AB"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
