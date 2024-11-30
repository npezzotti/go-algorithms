package largest_rectangle_in_histogram

type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:index]

	return elem
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}

	index := len(*s) - 1
	elem := (*s)[index]

	return elem
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func largestRectangleArea(heights []int) int {
	var maxArea int
	var stack Stack

	heights = append(heights, 0)

	for i, height := range heights {
		for !stack.IsEmpty() && heights[stack.Top()] > height {
			top := stack.Pop()

			width := i
			if !stack.IsEmpty() {
				width = i - stack.Top() - 1
			}

			maxArea = max(maxArea, heights[top]*width)
		}

		stack.Push(i)
	}

	return maxArea
}
