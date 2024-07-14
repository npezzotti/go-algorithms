package valid_parenthesis

import "testing"

func TestValidParenthesis(t *testing.T) {
	got := ValidParenthesis("()[]{}")
	want := true

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}

func TestValidParenthesisTwoChar(t *testing.T) {
	got := ValidParenthesis("(])")
	want := false

	if got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}
