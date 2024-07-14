package quick_sort

func QuickSort(nums []int, start int, end int) {
	if start < end {
		var pivot int = nums[end]
		var i int = start - 1

		for j := start; j < end; j++ {
			if nums[j] < pivot {
				i++
				nums[j], nums[i] = nums[i], nums[j]
			}
		}

		i++
		nums[i], nums[end] = nums[end], nums[i]

		QuickSort(nums, start, i-1)
		QuickSort(nums, i+1, end)
	}
}
