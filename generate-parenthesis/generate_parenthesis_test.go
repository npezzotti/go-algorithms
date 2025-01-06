package generate_parenthesis

import (
	"slices"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tcases := []struct {
		input  int
		output []string
	}{
		{
			input:  3,
			output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			input:  1,
			output: []string{"()"},
		},
	}

	for _, tt := range tcases {
		if output := generateParenthesis(tt.input); !slices.Equal(output, tt.output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
