package rotate_array

import (
	"slices"
)

func rotate(nums []int, k int) {
	pivot := k % len(nums)

	slices.Reverse(nums)
	slices.Reverse(nums[:pivot])
	slices.Reverse(nums[pivot:])
}
