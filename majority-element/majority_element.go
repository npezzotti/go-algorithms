package majority_element

func majorityElement(nums []int) int {
	var candidate, count = nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate, count = nums[i], 1
		} else {
			if candidate == nums[i] {
				count++
			} else {
				count--
			}
		}
	}

	return candidate
}
