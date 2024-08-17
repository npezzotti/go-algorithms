package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := map[string]struct {
		nums []int
		want [][]int
	}{
		"two triplets sum to zero": {
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		"one triplet does not sum to 0": {
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		"one triplet sums to 0": {
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			got := threeSum(tc.nums)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}
