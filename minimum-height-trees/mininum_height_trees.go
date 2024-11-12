package mininum_height_trees

func findMinHeightTrees(n int, edges [][]int) []int {
	if len(edges) == 0 {
		return []int{0}
	}

	adjMap := make(map[int][]int)
	leafCount := make([]int, n)

	for _, edge := range edges {
		adjMap[edge[0]] = append(adjMap[edge[0]], edge[1])
		leafCount[edge[0]]++
		adjMap[edge[1]] = append(adjMap[edge[1]], edge[0])
		leafCount[edge[1]]++
	}

	var minHeightTrees []int

	leaves := make([]int, 0)
	for idx, num := range leafCount {
		if num == 1 {
			leaves = append(leaves, idx)
		}
	}

	for len(leaves) > 0 {
		minHeightTrees = nil

		for _, leaf := range leaves {
			leaves = leaves[1:]
			minHeightTrees = append(minHeightTrees, leaf)

			for _, neighbor := range adjMap[leaf] {
				leafCount[neighbor]--

				if leafCount[neighbor] == 1 {
					leaves = append(leaves, neighbor)
				}
			}
		}
	}

	return minHeightTrees
}
