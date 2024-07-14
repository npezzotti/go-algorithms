package move_zeros

/* Given an integer array nums, move all 0's
to the end of it while maintaining the relative
order of the non-zero elements.

Note that you must do this in-place without
making a copy of the array.
*/

func moveZeroes(nums []int) {
	l := 0

	for r := range nums {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
