package coin_change

func coinChange(coins []int, amount int) int {
	var cache = make([]int, amount+1)

	for i := 1; i < amount+1; i++ {
		minCoins := amount + 1

		for _, coin := range coins {
			remain := i - coin

			if remain < 0 {
				continue
			}

			minCoins = min(minCoins, cache[remain]+1)
		}

		cache[i] = minCoins
	}

	if cache[amount] == amount+1 {
		return -1
	}

	return cache[amount]
}
