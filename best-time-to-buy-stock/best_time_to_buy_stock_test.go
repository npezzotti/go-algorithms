package best_time_to_buy_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	want := 5

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
