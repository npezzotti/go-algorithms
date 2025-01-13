package random_pick_with_weight

import (
	"math/rand"
)

type Solution struct {
	nums         []int
	prefixSumArr []int
}

func Constructor(w []int) Solution {
	prefixSumArr := make([]int, len(w)+1)
	prefixSumArr[0] = 0

	var curSum int
	for i, num := range w {
		curSum += num
		prefixSumArr[i+1] = curSum
	}

	return Solution{
		nums:         w,
		prefixSumArr: prefixSumArr,
	}
}

func (this *Solution) PickIndex() int {
	num := rand.Intn(this.prefixSumArr[len(this.prefixSumArr)-1]) + 1

	l, r := 0, len(this.prefixSumArr)-1
	for l < r {
		mid := (l + r) / 2

		if this.prefixSumArr[mid] >= num {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l-1
}
