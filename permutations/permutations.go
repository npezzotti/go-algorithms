package permutations

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var backTrack func(start int, curNums []int)
	backTrack = func(start int, curNums []int) {
		if start == len(nums) {
			temp := make([]int, len(curNums))
			copy(temp, curNums)
			res = append(res, temp)
		}

		for i := start; i < len(nums); i++ {
			curNums[start], curNums[i] = curNums[i], curNums[start]
			backTrack(start+1, curNums)
			curNums[i], curNums[start] = curNums[start], curNums[i]
		}
	}

	backTrack(0, nums)

	return res
}
