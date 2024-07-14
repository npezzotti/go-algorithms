package add_binary

import "testing"

func TestAddBinary(t *testing.T) {
	got := addBinary("11", "1")
	want := "100"

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
