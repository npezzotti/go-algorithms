package merge_sort

import "fmt"

func MergeSort(nums []int) []int {
	fmt.Println("merge sort called with", nums)
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	fmt.Println("merge called with", left, right)
	result := make([]int, 0, len(left)+len(right))

	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
