package roman_int_to_string

// Given a roman numeral, convert it to an integer.

var numerals = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	var total int = numerals[s[len(s)-1]]

	for i := 0; i < len(s)-1; i++ {
		if numerals[s[i]] < numerals[s[i+1]] {
			total -= numerals[s[i]]
		} else {
			total += numerals[s[i]]
		}
	}

	return total
}
