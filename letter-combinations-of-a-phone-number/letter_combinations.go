package letter

func letterCombinations(digits string) []string {
	res := make([]string, 0)

	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(curStr string, start int)
	backtrack = func(curStr string, start int) {
		if start == len(digits) {
			res = append(res, curStr)
			return
		}

		current_digit := digits[start]
		letters := digitMap[current_digit]

		for _, letter := range letters {
			backtrack(curStr+string(letter), start+1)
		}
	}

	if len(digits) > 0 {
		backtrack("", 0)
	}

	return res
}
