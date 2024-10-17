package atoi

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	var res, idx int
	var negative bool
	n := len(s)

	for idx < n && s[idx] == ' ' {
		idx++
	}

	if idx < n && (s[idx] == '+' || s[idx] == '-') {
		fmt.Println("plus or minus")
		if s[idx] == '-' {
			negative = true
		}
		idx++
	}

	for ; idx < n; idx++ {
		digit := int(s[idx] - '0')

		if digit < 0 || digit > 9 {
			break
		}

		res = res*10 + digit

		if res > math.MaxInt32 && !negative {
			return math.MaxInt32
		}

		if -res < math.MinInt32 && negative {
			return math.MinInt32
		}
	}

	if negative {
		res = -res
	}

	return res
}
