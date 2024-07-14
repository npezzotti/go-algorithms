package is_isomorphic

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can
 be replaced to get t.

All occurrences of a character must be replaced with another
character while preserving the order of characters. No two
characters may map to the same character, but a character may
map to itself.
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charMapS, charMapT = map[byte]byte{}, map[byte]byte{}

	for c := range s {
		if v, ok := charMapS[s[c]]; ok {
			if t[c] != v {
				return false
			}
		} else if _, ok := charMapT[t[c]]; ok {
			return false
		} else {
			charMapS[s[c]] = t[c]
			charMapT[t[c]] = s[c]
		}
	}

	return true
}
