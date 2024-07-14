package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	got := plusOne([]int{9, 9})
	want := []int{1, 0, 0}

	match := true
	if len(got) != len(want) {
		match = false
	}

	for idx := range got {
		if got[idx] != want[idx] {
			match = false
		}
	}

	if !match {
		t.Fatalf("got %v, want %v", got, want)
	}
}
