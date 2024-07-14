package length_of_last_word

func lengthOfLastWord(s string) int {
	var res int
	i := len(s) - 1

	for i >= 0 {
		if s[i] == ' ' {
			i--
		} else {
			break
		}
	}

	for i >= 0 {
		if s[i] == ' ' {
			break
		} else {
			i--
			res++
		}
	}
	
	return res
}
