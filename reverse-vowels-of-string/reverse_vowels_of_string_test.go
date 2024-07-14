package reverse_vowels_of_string

import "testing"

func TestReverseVowels(t *testing.T) {
	got := reverseVowels("golang")
	want := "galong"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
