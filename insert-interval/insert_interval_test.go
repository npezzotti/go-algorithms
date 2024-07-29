package insert_interval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tcases := []struct {
		name   string
		intArr [][]int
		newInt []int
		want   [][]int
	}{
		{
			name:   "non-overlapping",
			intArr: [][]int{{1, 3}, {6, 9}},
			newInt: []int{2, 5},
			want:   [][]int{{1, 5}, {6, 9}},
		},
		{
			name:   "overlapping",
			intArr: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInt: []int{4, 8},
			want:   [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			got := insert(tc.intArr, tc.newInt)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v\n", got, tc.want)
			}
		})
	}
}
