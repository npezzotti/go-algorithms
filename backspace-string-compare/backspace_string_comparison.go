package backspace_string_compare

func backspaceCompareStack(s string, t string) bool {
	return string(evalChars(s)) == string(evalChars(t))
}

func evalChars(s string) []rune {
	var stack []rune

	for _, char := range s {
		if char == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}

	return stack
}

func backspaceCompareTwoPointers(s string, t string) bool {
	nexValidIdx := func(str string, idx int) int {
		var backspaces int
		for idx >= 0 {
			if backspaces <= 0 && str[idx] != '#' {
				break
			} else if str[idx] == '#' {
				backspaces++
			} else {
				backspaces--
			}

			idx--
		}

		return idx
	}

	sIdx, tIdx := len(s)-1, len(t)-1
	for sIdx >= 0 || tIdx >= 0 {
		sIdx = nexValidIdx(s, sIdx)
		tIdx = nexValidIdx(t, tIdx)

		var sChar, tChar byte
		if sIdx != -1 {
			sChar = s[sIdx]
		}

		if tIdx != -1 {
			tChar = t[tIdx]
		}

		if sChar != tChar {
			return false
		}

		sIdx--
		tIdx--
	}

	return true
}
