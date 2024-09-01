package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	tcases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
	}

	for _, tc := range tcases {
		if got := coinChange(tc.coins, tc.amount); got != tc.want {
			t.Fatalf("got %d, want %d", got, tc.want)
		}
	}
}
