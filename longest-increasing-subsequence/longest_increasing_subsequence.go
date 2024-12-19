package longest_increasing_subsequence

import (
	"slices"
)

func lengthOfLISDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		for j := range i {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}

func lengthOfLISBinarySearch(nums []int) int {
	var lis []int

	for _, num := range nums {
		n := len(lis)
		left, right := 0, n
		for left < right {
			mid := (left + right) / 2
			if lis[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == n {
			lis = append(lis, num)
		} else {
			lis[left] = num
		}
	}

	return len(lis)
}
