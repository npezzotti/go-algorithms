package longest_consecutive_sequence

import (
	"sort"
)

func longestConsecutiveBruteForce(nums []int) int {
	sort.Ints(nums)

	var maxRes int
	curRes := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			curRes++
		} else {
			maxRes = max(maxRes, curRes)
			curRes = 1
		}
	}

	maxRes = max(maxRes, curRes)

	return maxRes
}

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	var longest int
	for num := range set {
		if _, ok := set[num-1]; !ok {
			start := num
			var curLongest int
			for {
				if _, ok := set[start]; !ok {
					break
				}

				start++
				curLongest++
			}

			longest = max(longest, curLongest)
		}
	}

	return longest
}
