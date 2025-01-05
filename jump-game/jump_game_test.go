package jump_game

import "testing"

func Test_canJump(t *testing.T) {
	tcases := []struct {
		nums   []int
		output bool
	}{
		{
			nums:   []int{2, 2, 1, 1, 4},
			output: true,
		},
		{
			nums:   []int{3, 2, 1, 0, 4},
			output: false,
		},
	}

	for _, tt := range tcases {
		if output := canJump(tt.nums); output != tt.output {
			t.Errorf("got %t, expected %t", output, tt.output)
		}
	}
}
