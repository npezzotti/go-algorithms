package house_robber

func rob(nums []int) int {
	n := len(nums)

	cache := make(map[int]int)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	cache[0], cache[1] = nums[0], max(nums[0], nums[1])
	var backtrack func(i int) int
	backtrack = func(i int) int {
		if n, ok := cache[i]; ok {
			return n
		}

		cache[i] = max(nums[i]+backtrack(i-2), backtrack(i-1))
		return cache[i]
	}

	return backtrack(n - 1)
}
