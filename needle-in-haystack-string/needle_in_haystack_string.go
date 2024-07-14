package needle_in_haystack_string

func NeedleInHaystackString(haystack string, needle string) int {
	idx := -1
	for i := 0; i <= (len(haystack) - len(needle)); i++ {
		idx = i
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				idx = -1
				break
			}
		}

		if idx != -1 {
			return idx
		}
	}

	return idx
}
