package largest_number

import "testing"

func Test_largestNumber(t *testing.T) {
	tcases := []struct {
		nums   []int
		output string
	}{
		{
			nums:   []int{10, 2},
			output: "210",
		},
		{
			nums:   []int{3, 30, 34, 5, 9},
			output: "9534330",
		},
		{
			nums:   []int{111311, 1113},
			output: "1113111311",
		},
		{
			nums:   []int{34323, 3432},
			output: "343234323",
		},
		{
			nums:   []int{0, 0},
			output: "0",
		},
	}

	for _, tt := range tcases {
		if output := largestNumber(tt.nums); output != tt.output {
			t.Errorf("got %s, expected %s", output, tt.output)
		}
	}
}
