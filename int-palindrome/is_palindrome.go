package is_palindrome

/*
Given an integer x, return true if x
is a palindrome, and false otherwise.
*/

func IsPalindrome(s string) bool {
	length := len(s)

	if length < 2 {
		return true
	}

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}
