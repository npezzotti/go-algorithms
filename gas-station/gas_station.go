package gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	var totalGas, currentGas, startStation int
	for i := 0; i < len(gas); i++ {
		totalGas += (gas[i] - cost[i])
		currentGas += (gas[i] - cost[i])

		if currentGas < 0 {
			startStation = i + 1
			currentGas = 0
		}
	}

	if totalGas < 0 {
		startStation = -1
	}

	return startStation
}
