package three_sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	var res int = math.MaxInt
	for idx1 := 0; idx1 < n-2; idx1++ {
		idx2, idx3 := idx1+1, n-1
		for idx2 < idx3 {
			sum := nums[idx1] + nums[idx2] + nums[idx3]

			if sum == target {
				return sum
			}

			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}

			if sum < target {
				idx2++
			} else {
				idx3--
			}
		}
	}

	return res
}
