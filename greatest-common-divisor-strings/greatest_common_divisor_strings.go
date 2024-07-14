package greatest_common_divisor_strings

import "fmt"

/*
For two strings s and t, we say "t divides s" if
and only if s = t + ... + t (i.e., t is concatenated
with itself one or more times).

Given two strings str1 and str2, return the largest
string x such that x divides both str1 and str2.
*/

func gcdOfStrings(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	if str1[:len(str2)] != str2 {
		return ""
	}

	return gcdOfStrings(str1[len(str2):], str2)
}
