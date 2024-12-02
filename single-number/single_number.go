package single_number

func singleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res = res ^ num
	}

	return res
}
