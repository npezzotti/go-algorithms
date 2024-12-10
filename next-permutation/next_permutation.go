package next_permutation

import (
	"slices"
)

func nextPermutation(nums []int) {
	n := len(nums)
	pivot := n - 2

	for pivot >= 0 && nums[pivot] >= nums[pivot+1] {
		pivot--
	}

	if pivot == -1 {
		slices.Reverse(nums)
		return
	}

	for i := n - 1; i > pivot; i-- {
		if nums[i] > nums[pivot] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			break
		}
	}

	l, r := pivot+1, n-1

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
