package maximal_square

import "testing"

func Test_maximalSquare(t *testing.T) {
	tcases := []struct {
		matrix [][]byte
		output int
	}{
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'}},
			output: 4,
		},
		{
			matrix: [][]byte{
				{'0', '1'},
				{'1', '0'}},
			output: 1,
		},
		{
			matrix: [][]byte{{0}},
			output: 0,
		},
	}

	for _, tt := range tcases {
		if output := maximalSquare(tt.matrix); output != tt.output {
			t.Errorf("got %d, expected %d", output, tt.output)
		}
	}
}
