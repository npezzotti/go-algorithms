package daily_temperatures

type Stack []int

func (s *Stack) Pop() (x int) {
	x = s.Peek()
	*s = (*s)[:len(*s)-1]
	return
}

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) Peek() (x int) {
	x = (*s)[len(*s)-1]
	return
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	var s Stack
	for idx, temp := range temperatures {
		for !s.Empty() && temperatures[s.Peek()] < temp {
			tIdx := s.Pop()
			res[tIdx] = idx - tIdx
		}

		s.Push(idx)
	}

	return res
}
