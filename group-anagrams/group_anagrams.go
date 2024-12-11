package group_anagrams

type AnagramHashKey [26]byte

func groupAnagrams(strs []string) [][]string {
	groups := make(map[AnagramHashKey][]string)

	for _, str := range strs {
		var k AnagramHashKey
		for _, char := range str {
			k[char-'a']++
		}

		groups[k] = append(groups[k], str)
	}

	var res [][]string
	for _, group := range groups {
		res = append(res, group)
	}

	return res
}
