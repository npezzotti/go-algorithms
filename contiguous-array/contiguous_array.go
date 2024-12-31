package contiguous_array

func findMaxLengthBrute(nums []int) int {
	var count int
	for idx := range nums {
		var zeros, ones int
		for j := idx; j < len(nums); j++ {
			if nums[j] == 0 {
				zeros++
			} else {
				ones++
			}

			if zeros == ones {
				count = max(count, j+1-idx)
			}
		}
	}

	return count
}

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
