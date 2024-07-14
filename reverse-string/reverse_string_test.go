package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte("hello")
	reverseString(s)

	want := []byte("olleh")

	if string(s) != string(want) {
		t.Fatalf("got %s, want %s", s, want)
	}
}
