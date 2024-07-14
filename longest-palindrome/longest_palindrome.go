package longest_palindrome

func longestPalindrome(s string) int {
	numMap := make(map[rune]int)

	var pairs int
	for _, strChar := range s {
		numMap[strChar]++

		if numMap[strChar]%2 == 0 {
			pairs++
		}
	}

	if pairs*2 == len(s) {
		return len(s)
	}

	return pairs*2 + 1
}
