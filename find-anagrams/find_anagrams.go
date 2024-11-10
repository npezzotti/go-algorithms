package find_anagrams

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	if len(s) < len(p) {
		return res
	}

	pCount := make([]int, 26)
	windowCount := make([]int, 26)

	for _, char := range p {
		pCount[char-'a']++
	}

	for _, char := range s[0:len(p)] {
		windowCount[char-'a']++
	}

	if arraysEqual(pCount, windowCount) {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		windowCount[s[i]-'a']++
		windowCount[s[i-len(p)]-'a']--

		if arraysEqual(pCount, windowCount) {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}

func arraysEqual(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
