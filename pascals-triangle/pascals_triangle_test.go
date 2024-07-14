package pascals_triangle

import "testing"

func TestGenerate(t *testing.T) {
	got := generate(5)
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}

	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("got %v, want %v", got, want)
		}

		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Fatalf("got %v, want %v", got, want)
			}
		}
	}
}
