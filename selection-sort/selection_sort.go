package selection_sort

func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := nums[i] + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
