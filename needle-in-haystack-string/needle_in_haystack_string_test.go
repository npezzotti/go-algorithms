package needle_in_haystack_string

import "testing"

func TestNeedleInHaystackStringFail(t *testing.T) {
	got := NeedleInHaystackString("jjjjjjjjs", "sad")
	want := -1

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestNeedleInHaystackStringPass(t *testing.T) {
	got := NeedleInHaystackString("sadbutsad", "sad")
	want := 0

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
