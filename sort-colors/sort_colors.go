package sort_colors

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	current := 0

	for current <= right {
		if nums[current] == 0 {
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		} else if nums[current] == 2 {
			nums[right], nums[current] = nums[current], nums[right]
			right--
		} else{
			current++
		}
	}
}
