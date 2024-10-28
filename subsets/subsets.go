package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(0)

	return res
}
