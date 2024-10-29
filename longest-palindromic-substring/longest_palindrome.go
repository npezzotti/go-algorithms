package longest_palindrome

func longestPalindrome(s string) string {
	var resLen, start, end int

	for i := range s {
		maxLen := max(expandFromCenter(i, i, s), expandFromCenter(i, i+1, s))
		if maxLen > resLen {
			resLen = maxLen
			start = i - (resLen-1)/2
			end = i + resLen/2
		}
	}

	return string(s[start : end+1])
}

func expandFromCenter(l, r int, s string) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}
