package valid_parenthesis

/*
Given a string s containing just the characters '(', ')', '{', '}', 
'[' and ']', determine if the input string is valid. An input 
string is valid if:

	1. Open brackets must be closed by the same type of brackets.
	2. Open brackets must be closed in the correct order.
	3. Every close bracket has a corresponding open bracket of the same type.
*/

func ValidParenthesis(s string) bool {
	str := []byte(s)

	charMap := map[byte]byte{
		'[': ']',
		'{': '}',
		'(': ')',
	}

	var nextChars []byte
	for _, c := range str {

		if closing, ok := charMap[c]; ok {
			nextChars = append(nextChars, closing)
			continue
		}

		if len(nextChars) == 0 {
			return false
		}

		if nextChars[len(nextChars)-1] != c {
			return false
		}

		nextChars = nextChars[:len(nextChars)-1]
	}

	return len(nextChars) == 0
}
