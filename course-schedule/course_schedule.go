package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, prereq := range prerequisites {
		x, y := prereq[0], prereq[1]

		graph[x] = append(graph[x], y)
	}

	visited := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if hasCycle(graph, visited, i) {
			return false
		}
	}

	return true
}

func hasCycle(graph [][]int, visited []int, node int) bool {
	if visited[node] == 1 {
		return true
	}

	if visited[node] == 2 {
		return false
	}

	visited[node] = 1

	for _, edge := range graph[node] {
		if hasCycle(graph, visited, edge) {
			return true
		}
	}

	visited[node] = 2
	return false
}
