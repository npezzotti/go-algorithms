package subarry_sum_equals_k

func subarraySumBruteForce(nums []int, k int) int {
	var res int
	n := len(nums)

	for i := range n {
		var curSum int
		for j := i; j < n; j++ {
			curSum += nums[j]

			if curSum == k {
				res++
			}
		}
	}

	return res
}

func subarraySum(nums []int, k int) int {
	var res int
	
	var curSum int
	prefixSums := map[int]int{0: 1}
	for _, num := range nums {
		curSum += num

		if count, ok := prefixSums[curSum-k]; ok {
			res += count
		}

		prefixSums[curSum]++
	}

	return res
}
