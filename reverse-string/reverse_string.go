package reverse_string

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
}
