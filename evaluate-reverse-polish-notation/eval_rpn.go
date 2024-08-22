package eval_rpn

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		} else {
			if num, err := strconv.Atoi(token); err == nil {
				stack = append(stack, num)
			}
		}
	}
	return stack[len(stack)-1]
}
