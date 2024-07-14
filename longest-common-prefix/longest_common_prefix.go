package longest_common_prefix

import (
	"strings"
)

/*
Write a function to find the longest common prefix
string amongst an array of strings. If there is no
common prefix, return an empty string "".
*/

func longestCommonPrefix(strs []string) string {
	pref := strs[0]
	for _, s := range strs {
		for !strings.HasPrefix(s, pref) {
			pref = pref[:len(pref)-1]
		}
	}
	return pref
}
