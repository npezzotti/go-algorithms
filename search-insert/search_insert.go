package search_insert

func SearchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums)
}
