package repeated_substring_pattern

import "strings"

/* Given a string s, check if it can be constructed
by taking a substring of it and appending multiple
copies of the substring together.
*/

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			substr := s[:i]

			var res []byte
			for j := 0; j < len(s)/i; j++ {
				res = append(res, substr...)
			}

			if string(res) == s {
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPatternStandardLib(s string) bool {
	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)
}
