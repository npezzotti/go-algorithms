package reverse_integer

import (
	"math"
)

func reverse(x int) int {
	var res int

	for x > 9 || x < -9 {
		if res > math.MaxInt32/10 || res < math.MinInt32/10 {
			return 0
		}

		res = res * 10
		a := x % 10
		res += a * 10
		x = (x - a) / 10
	}

	res += x

	return res
}
