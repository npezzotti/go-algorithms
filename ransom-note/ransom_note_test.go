package ransom_note

import "testing"

func TestCanConstruct(t *testing.T) {
	got := canConstruct("hello", "loleh")
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
