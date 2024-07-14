package is_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	got := isPalindrome("A man, a plan, a canal: Panama")
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
