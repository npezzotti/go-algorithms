package trapping_rain_water

func trapBruteForce(height []int) int {
	var trappedWater int

	for i := 1; i < len(height); i++ {
		var leftMax int
		for j := i; j >= 0; j-- {
			leftMax = max(height[j], leftMax)
		}

		var rightMax int
		for j := i; j < len(height); j++ {
			rightMax = max(height[j], rightMax)
		}

		trappedWater += max(min(leftMax, rightMax)-height[i], 0)
	}

	return trappedWater
}

func trapBruteForceTwoArrs(height []int) int {
	var trappedWater int
	n := len(height)

	leftMaxArr, rightMaxArr := make([]int, n), make([]int, n)
	var leftMax, rightMax int
	for i := 0; i < len(height); i++ {
		leftMax = max(height[i], leftMax)
		leftMaxArr[i] = leftMax
	}

	for i := len(height) - 1; i >= 0; i-- {
		rightMax = max(height[i], rightMax)
		rightMaxArr[i] = rightMax
	}

	for i := 0; i < len(height); i++ {
		trappedWater += max(min(leftMaxArr[i], rightMaxArr[i])-height[i], 0)
	}

	return trappedWater
}

type Stack []int

func (s *Stack) Push(num int) {
	*s = append((*s), num)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	num := s.Top()
	*s = (*s)[:len(*s)-1]

	return num
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}

	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func trapStack(height []int) int {
	var trappedWater int

	s := new(Stack)
	var i int
	for i < len(height) {
		for !s.IsEmpty() && height[i] > height[s.Top()] {
			topTower := s.Pop()

			if s.IsEmpty() {
				break
			}

			areaWidth := i - s.Top() - 1
			areaHeight := min(height[i], height[s.Top()]) - height[topTower]

			trappedWater += areaWidth * areaHeight
		}

		s.Push(i)
		i++
	}

	return trappedWater
}

func trapTwoPointers(height []int) int {
	var trappedWater int
	var maxLeft, maxRight int

	for l, r := 0, len(height)-1; l <= r; {
		if height[l] < height[r] {
			if height[l] > maxLeft {
				maxLeft = height[l]
			} else {
				trappedWater += maxLeft - height[l]
			}
			l++
		} else {
			if height[r] > maxRight {
				maxRight = height[r]
			} else {
				trappedWater += maxRight - height[r]
			}
			r--
		}
	}

	return trappedWater
}
