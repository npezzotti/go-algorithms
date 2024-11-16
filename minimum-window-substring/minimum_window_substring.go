package minimum_window_substring

func minWindow(s string, t string) string {
	var res string

	var rem int
	charMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		rem++
		charMap[t[i]]++
	}

	for l, r := 0, 0; r < len(s); r++ {
		if cnt, ok := charMap[s[r]]; ok {
			// check for duplicates
			if cnt > 0 {
				rem--
			}
			charMap[s[r]]--
		}

		for rem == 0 {
			if res == "" || len(res) > len(s[l:r+1]) {
				res = s[l : r+1]
			}

			if cnt, ok := charMap[s[l]]; ok {
				if cnt >= 0 {
					rem++
				}
				charMap[s[l]]++
			}

			l++
		}
	}

	return res
}
