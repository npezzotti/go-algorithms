package binary_search

/*
Given an array of integers nums which is sorted in ascending
order, and an integer target, write a function to search
target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0

	for low <= high {
		mid := (high + low) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
