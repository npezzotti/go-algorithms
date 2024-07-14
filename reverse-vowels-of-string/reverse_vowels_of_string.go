package reverse_vowels_of_string

import (
	"bytes"
)

func reverseVowels(s string) string {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	byteString := []byte(s)

	for i, j := 0, len(byteString)-1; i <= j; {
		isIVowel := bytes.IndexByte(vowels, byteString[i]) != -1
		isJVowel := bytes.IndexByte(vowels, byteString[j]) != -1

		if isIVowel && isJVowel {
			byteString[i], byteString[j] = byteString[j], byteString[i]
		}

		if !isIVowel {
			i++
		}

		if !isJVowel {
			j--
		}
	}

	return string(byteString)
}
