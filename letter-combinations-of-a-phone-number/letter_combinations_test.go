package letter

import (
	"slices"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	tcases := []struct {
		digits string
		result []string
	}{
		{
			digits: "23",
			result: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			result: []string{},
		},
		{
			digits: "2",
			result: []string{"a", "b", "c"},
		},
	}

	for _, tc := range tcases {
		if result := letterCombinations(tc.digits); !slices.Equal(result, tc.result) {
			t.Errorf("expected %+v, got %+v", tc.result, result)
		}
	}
}
