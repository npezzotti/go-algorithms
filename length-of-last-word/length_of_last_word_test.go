package length_of_last_word

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	got := lengthOfLastWord("  fly me to the moon  ")
	want := 4

	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}
}
