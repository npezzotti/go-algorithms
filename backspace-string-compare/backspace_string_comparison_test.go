package backspace_string_compare

import (
	"fmt"
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	tcases := []struct {
		s, t   string
		output bool
	}{
		{
			s:      "ab#c",
			t:      "ad#c",
			output: true,
		},
		{
			s:      "ab##",
			t:      "c#d#",
			output: true,
		},
		{
			s:      "a#c",
			t:      "b",
			output: false,
		},
		{
			s:      "bbbextm",
			t:      "bbb#extm",
			output: false,
		},
	}

	for _, tc := range tcases {
		t.Run(fmt.Sprintf("backspaceCompareStack:%s,%s", tc.s, tc.t), func(t *testing.T) {
			if output := backspaceCompareStack(tc.s, tc.t); output != tc.output {
				t.Errorf("got %t, expected %t", output, tc.output)
			}
		})

		t.Run(fmt.Sprintf("backspaceCompareTwoPointers:%s,%s", tc.s, tc.t), func(t *testing.T) {
			if output := backspaceCompareTwoPointers(tc.s, tc.t); output != tc.output {
				t.Errorf("got %t, expected %t", output, tc.output)
			}
		})
	}
}
