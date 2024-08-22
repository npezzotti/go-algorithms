package eval_rpn

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tcases := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}
	for _, tc := range tcases {
		t.Run(fmt.Sprintf("%v", tc.tokens), func(t *testing.T) {
			got := evalRPN(tc.tokens)
			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}
