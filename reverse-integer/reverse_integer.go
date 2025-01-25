package reverse_integer

import (
	"math"
)

func reverse(x int) int {
	var res int

	for x != 0 {
		ones := x % 10
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}

		res = res*10 + ones
		x /= 10
	}

	return res
}
