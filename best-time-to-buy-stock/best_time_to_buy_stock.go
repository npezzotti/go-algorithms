package best_time_to_buy_stock

/*
You are given an array prices where prices[i] is the
price of a given stock on the ith day. You want to
maximize your profit by choosing a single day to buy
one stock and choosing a different day in the future
to sell that stock.

Return the maximum profit you can achieve from this
transaction. If you cannot achieve any profit, return 0.
*/

func maxProfit(prices []int) int {
	var max int
	low := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		} else if prices[i]-low > max {
			max = prices[i] - low
		}
	}

	return max
}
