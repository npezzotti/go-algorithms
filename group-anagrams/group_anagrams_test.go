package group_anagrams

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tcases := []struct {
		strs   []string
		output [][]string
	}{
		{
			strs:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			strs:   []string{""},
			output: [][]string{{""}},
		},
		{
			strs:   []string{"a"},
			output: [][]string{{"a"}},
		},
	}

	for _, tc := range tcases {
		if res := groupAnagrams(tc.strs); !areStringSlicesEqual(res, tc.output) {
			t.Errorf("got %+v, expected %+v", res, tc.output)
		}
	}
}

func areStringSlicesEqual(strSlice1, strSlice2 [][]string) bool {
	if len(strSlice1) != len(strSlice2) {
		return false
	}

	for i := range strSlice1 {
		if len(strSlice1[i]) != len(strSlice2[i]) {
			return false
		}

		for j := range strSlice1[i] {
			if strSlice1[i][j] != strSlice2[i][j] {
				return false
			}
		}
	}

	return true
}
