package basic_calculator

import (
	"unicode"
)

func calculate(s string) int {
	cur := 0
	sign := 1
	exprLen := len(s)

	stack := make([]int, 0)
	for i := 0; i < exprLen; i++ {
		char := s[i]
		switch char {
		case ' ':
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, cur, sign)
			cur = 0
			sign = 1
		case ')':
			s := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			cur = c + cur*s
		default:
			var n int
			start := i
			for start < exprLen && unicode.IsDigit(rune(s[start])) {
				n = n*10 + int(s[start]-'0')
				start++
			}

			cur = cur + n*sign
			i = start - 1
		}
	}

	return cur
}
