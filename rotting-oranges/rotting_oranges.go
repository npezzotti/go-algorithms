package rotting_oranges

const (
	Empty = iota
	Fresh
	Rotten
)

type coordinate = [2]int

type coordinates []coordinate

func orangesRotting(grid [][]int) int {
	var fresh int
	rows, cols := len(grid), len(grid[0])

	rotten := make(coordinates, 0)
	for r := range grid {
		for c := range grid[r] {
			switch grid[r][c] {
			case Fresh:
				fresh++
			case Rotten:
				rotten = append(rotten, coordinate{r, c})
			default:
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	var minutes int
	var directions = coordinates{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(rotten) > 0 {
		for _, coord := range rotten {
			for _, direction := range directions {
				adjRow, adjCol := coord[0]+direction[0], coord[1]+direction[1]
				if adjRow >= 0 && adjRow < rows && adjCol >= 0 && adjCol < cols && grid[adjRow][adjCol] == Fresh {
					grid[adjRow][adjCol] = Rotten
					rotten = append(rotten, coordinate{adjRow, adjCol})
					fresh--
				}
			}
			rotten = rotten[1:]
		}
		minutes++

		if fresh == 0 {
			return minutes
		}
	}

	return -1
}
