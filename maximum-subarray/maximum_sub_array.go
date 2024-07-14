package maximum_sub_array

func maxSubArray(nums []int) int {
	var max, cur = nums[0], nums[0]

	for _, num := range nums[1:] {
		if cur < 0 {
			cur = 0
		}

		cur += num

		if max < cur {
			max = cur
		}
	}

	return max
}
