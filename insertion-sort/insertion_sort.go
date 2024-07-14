package insertion_sort

func InsertionSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
