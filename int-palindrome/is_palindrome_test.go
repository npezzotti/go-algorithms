package is_palindrome

import (
	"testing"
)

func TestIsPalindromePass(t *testing.T) {
	got := IsPalindrome("tacocat")
	want := true

	if !got {
		t.Fatalf("expected %t, got %t", want, got)
	}
}

func TestIsPalindromeFail(t *testing.T) {
	got := IsPalindrome("hello")
	want := false

	if got {
		t.Fatalf("expected %t, got %t", want, got)
	}
}

func TestIsPalindromeOneCharStringPass(t *testing.T) {
	got := IsPalindrome("a")
	want := true

	if !got {
		t.Fatalf("expected %t, got %t", want, got)
	}
}
