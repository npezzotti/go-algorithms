package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	length := len(nums)

	if length < 2 {
		return length
	}

	p := 0
	for i := 1; i < length; i++ {
		if nums[p] != nums[i] {
			p++
			nums[p] = nums[i]
		}
	}

	return p + 1
}
