package can_partition

func canPartition(nums []int) bool {
	var total int
	for _, i := range nums {
		total += i
	}

	if total%2 != 0 {
		return false
	}

	targetSum := total / 2

	dp := make([]bool, targetSum)
	dp[0] = true
	for _, num := range nums {
		// ignore numbers greater than targetSum
		if num <= targetSum {
			// found compliment to targetSum, which
			// means remaining nums sum to targetSum
			if dp[targetSum-num] {
				return true
			}
			// loop backwards, starting at (targetNum - n - 1), since
			// we already confirmed above that (targetNum - n) is not true
			for j := targetSum - num - 1; j >= 0; j-- {
				if dp[j] {
					dp[j+num] = true
				}
			}
		}
	}
	return false
}
