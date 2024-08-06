package length_of_longest_substring

func lengthOfLongestSubstring(s string) int {
	res := 0
	seen := make([]int, 256)

	for l, r := 0, 0; r < len(s); r++ {
		seen[s[r]]++
		for seen[s[r]] > 1 {
			seen[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
