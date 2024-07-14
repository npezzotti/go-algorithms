package is_anagram

func IsAnagaram(s, t string) bool {
	charArr := make([]int, 26)

	for _, char := range s {
		i := int(char - 'a')
		charArr[i]++
	}

	for _, char := range t {
		i := int(char - 'a')
		charArr[i]--
	}

	for _, count := range charArr {
		if count != 0 {
			return false
		}
	}

	return true
}
