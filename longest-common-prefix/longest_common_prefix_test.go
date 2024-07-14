package longest_common_prefix

import "testing"

func TestLogestCommonPrefixMultipleChars(t *testing.T) {
	got := longestCommonPrefix([]string{
		"flower", "flow", "flight",
	})
	want := "fl"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestLogestCommonPrefixSingleChar(t *testing.T) {
	got := longestCommonPrefix([]string{
		"hello", "heat", "h",
	})
	want := "h"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
