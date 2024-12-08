package house_robber

func rob(nums []int) int {
	n := len(nums)

	var backtrack func(i int) int
	backtrack = func(i int) int {
		if i == 0 {
			return nums[0]
		}

		if i == 1 {
			return max(nums[0], nums[1])
		}

		return max(nums[i]+backtrack(i-2), backtrack(i-1))
	}

	return backtrack(n - 1)
}
