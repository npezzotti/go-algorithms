package repeated_substring_pattern

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	got := repeatedSubstringPattern("abcabc")
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}

func TestRepeatedSubstringPatternStandardLib(t *testing.T) {
	got := repeatedSubstringPatternStandardLib("abcdabcd")
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
