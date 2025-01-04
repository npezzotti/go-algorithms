package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	n := len(s)

	var maxLen int
	charCount := make([]int, 26)
	var l, r int
	for ; r < n; r++ {
		charCount[s[r]-'A']++
		maxLen = max(maxLen, charCount[s[r]-'A'])

		if r-l+1 > maxLen+k {
			charCount[s[l]-'A']--
			l++
		}
	}

	return r - l
}
