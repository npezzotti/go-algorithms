package decode_string

import (
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var stack []rune
	var count int
	var subStr string

	for _, char := range s {
		if char != ']' {
			stack = append(stack, char)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				subStr = string(stack[len(stack)-1]) + subStr
				stack = stack[:len(stack)-1]
			}

			stack = stack[:len(stack)-1]

			multiplier := 1
			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				count += multiplier * int(stack[len(stack)-1]-'0')
				stack = stack[:len(stack)-1]
				multiplier *= 10
			}

			subStr = strings.Repeat(subStr, count)
			stack = append(stack, []rune(subStr)...)

			count = 0
			subStr = ""
		}
	}

	return string(stack)
}
