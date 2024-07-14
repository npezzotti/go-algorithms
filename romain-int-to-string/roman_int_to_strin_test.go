package roman_int_to_string

import "testing"

func TestRomanToIntTwelve(t *testing.T) {
	got := RomanToInt("XII")
	want := 12

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestRomanToIntOneThousandNineHundredNinetyFour(t *testing.T) {
	got := RomanToInt("MCMXCIV")
	want := 1994

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
