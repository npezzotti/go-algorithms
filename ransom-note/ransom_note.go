package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	charMap := make([]int, 26)

	for _, c := range magazine {
		charMap[c-'a']++
	}

	for _, c := range ransomNote {
		charMap[c-'a']--

		if charMap[c-'a'] < 0 {
			return false
		}
	}

	return true
}
