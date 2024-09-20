package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	var backTrack func(idx, curSum int, curNums []int)
	backTrack = func(curIdx, curSum int, curNums []int) {
		if curSum == target {
			temp := make([]int, len(curNums))
			copy(temp, curNums)
			res = append(res, temp)
		} else if curSum > target {
			return
		} else {
			for i := curIdx; i < len(candidates); i++ {
				curNums = append(curNums, candidates[i])
				backTrack(i, curSum+candidates[i], curNums)
				curNums = curNums[:len(curNums)-1]
			}
		}
	}

	backTrack(0, 0, []int{})
	
	return res
}
