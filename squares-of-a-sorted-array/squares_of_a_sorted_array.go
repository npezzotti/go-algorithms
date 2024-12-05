package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	resIdx := len(res) - 1

	i, j := 0, n-1
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[resIdx] = (nums[i] * nums[i])
			i++
		} else {
			res[resIdx] = (nums[j] * nums[j])
			j--
		}

		resIdx--
	}

	return res
}
