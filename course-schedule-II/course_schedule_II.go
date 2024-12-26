package course_schedule_II

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, prerequisite := range prerequisites {
		x, y := prerequisite[0], prerequisite[1]
		graph[x] = append(graph[x], y)
	}

	visited := make([]int, numCourses)

	var res []int
	for i := range numCourses {
		if dfs(graph, visited, i, &res) {
			return []int{}
		}
	}

	return res
}

func dfs(graph [][]int, visited []int, node int, res *[]int) bool {
	if visited[node] == 1 {
		return true
	}

	if visited[node] == 2 {
		return false
	}

	visited[node] = 1

	for _, edge := range graph[node] {
		if dfs(graph, visited, edge, res) {
			return true
		}
	}

	*res = append(*res, node)
	visited[node] = 2
	return false
}
