package contiguous_array

func findMaxLength(nums []int) int {
	var ans, count int
	seen := make(map[int]int)
	seen[0] = -1

	for idx, num := range nums {
		if num == 0 {
			count--
		} else {
			count++
		}

		if prevIdx, ok := seen[count]; ok {
			ans = max(ans, idx-prevIdx)
		} else {
			seen[count] = idx
		}
	}

	return ans
}
