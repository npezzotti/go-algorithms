package container_with_most_water

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	var maxWater int

	for l < r {
		var mw int
		if height[l] < height[r] {
			mw = height[l] * (r - l)
			l++
		} else {
			mw = height[r] * (r - l)
			r--
		}

		if mw > maxWater {
			maxWater = mw
		}
	}

	return maxWater
}
