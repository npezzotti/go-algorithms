package missing_number

func missingNumberHash(nums []int) int {
	arr := make([]int, len(nums)+1)
	for _, num := range nums {
		arr[num]++
	}

	var res int
	for idx, count := range arr {
		if count == 0 {
			res = idx
		}
	}

	return res
}

func missingNumberSum(nums []int) int {
	n := len(nums)

	expectedSum := n * (n + 1) / 2

	var sum int
	for _, num := range nums {
		sum += num
	}

	return expectedSum - sum
}

func missingNumberBitwise(nums []int) int {
	missing := len(nums)

	for i, num := range nums {
		missing ^= i ^ num
	}

	return missing
}
