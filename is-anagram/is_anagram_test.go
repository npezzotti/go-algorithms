package is_anagram

import "testing"

func TestIsAnagramPass(t *testing.T) {
	res := IsAnagaram("anagram", "nagaram")

	if !res {
		t.Fatalf("got %v, want %v", res, true)
	}
}

func TestIsAnagramFail(t *testing.T) {
	res := IsAnagaram("rat", "car")

	if res {
		t.Fatalf("got %v, want %v", res, false)
	}
}
