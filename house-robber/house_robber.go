package house_robber

func robTopDownMemoized(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	cache := make(map[int]int)
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

func robBottomUpDP(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+max(dp[i-2]), dp[i-1])
	}

	return dp[n-1]
}
