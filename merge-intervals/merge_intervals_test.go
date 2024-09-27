package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tcases := []struct {
		intervals [][]int
		result    [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			result:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			result:    [][]int{{1, 5}},
		},
	}

	for _, tc := range tcases {
		if res := merge(tc.intervals); !reflect.DeepEqual(res, tc.result) {
			t.Fatalf("got %v, want %v", res, tc.result)
		}
	}
}
