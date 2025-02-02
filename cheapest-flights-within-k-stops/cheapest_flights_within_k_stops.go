package cheapest_flights_within_k_stops_test

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	distances := make([]int, n)
	for i := range distances {
		distances[i] = math.MaxInt32
	}

	distances[src] = 0

	for i := 0; i < k+1; i++ {
		distancesBackup := make([]int, n)
		copy(distancesBackup, distances)

		for _, flight := range flights {
			from, to, cost := flight[0], flight[1], flight[2]
			distances[to] = min(distances[to], distancesBackup[from]+cost)
		}
	}

	if distances[dst] == math.MaxInt32 {
		return -1
	}

	return distances[dst]
}
