package basic_calculator_iii

import (
	"unicode"
)

func calculate(s string) int {
	var cur int
	var prevSign rune = '+'

	stack := make([]int, 0)
	for idx, char := range s {
		if unicode.IsDigit(char) {
			cur = cur*10 + int(char-'0')
		}

		if char == '+' || char == '-' || char == '*' || char == '/' || idx == len(s)-1 {
			switch prevSign {
			case '+':
				stack = append(stack, cur)
			case '-':
				stack = append(stack, -cur)
			case '*':
				lastNum := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, cur*lastNum)
			case '/':
				lastNum := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, lastNum/cur)
			}

			prevSign = char
			cur = 0
		}
	}

	for len(stack) > 0 {
		num := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		cur += num
	}

	return cur
}
