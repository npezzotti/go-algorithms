package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	for idx, interval := range intervals {
		if newInterval[1] < interval[0] {
			result = append(result, newInterval)
			return append(result, intervals[idx:]...)
		} else if newInterval[0] > interval[1] {
			result = append(result, interval)
		} else {
			newInterval = []int{
				min(newInterval[0], interval[0]),
				max(newInterval[1], interval[1]),
			}
		}
	}

	result = append(result, newInterval)

	return result
}
