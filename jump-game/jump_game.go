package jump_game

func canJump(nums []int) bool {
	var maxIdx int
	for i := range nums {
		if maxIdx < i {
			return false
		}

		maxIdx = max(maxIdx, nums[i]+i)
	}

	return true
}
