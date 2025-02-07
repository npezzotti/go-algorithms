package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	var res = make([][]int, 0)
	for idx1 := 0; idx1 < n-2; idx1++ {
		if idx1 > 0 && nums[idx1] == nums[idx1-1] {
			continue
		}

		idx2 := idx1 + 1
		idx3 := n - 1
		for idx2 < idx3 {
			sum := nums[idx1] + nums[idx2] + nums[idx3]
			if sum == 0 {
				res = append(res, []int{nums[idx1], nums[idx2], nums[idx3]})

				tmpIdx3 := nums[idx3]
				for idx3 > idx2 && nums[idx3] == tmpIdx3 {
					idx3--
				}
			} else if sum > 0 {
				idx3--
			} else {
				idx2++
			}
		}
	}

	return res
}
