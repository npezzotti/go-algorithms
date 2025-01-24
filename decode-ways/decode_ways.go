package decode_ways

import "fmt"

func numDecodings(s string) int {
	prev, cur := 0, 1

	for idx, num := range s {
		fmt.Println(prev, cur)
		var h int
		if num != '0' {
			h = cur
		}

		if idx > 0 && s[idx-1] != '0' {
			prevNum := int(s[idx-1]-'0')*10 + int(s[idx]-'0')
			if prevNum <= 26 {
				h += prev
			}
		}

		prev = cur
		cur = h
	}
	fmt.Println(prev, cur)

	return cur
}
