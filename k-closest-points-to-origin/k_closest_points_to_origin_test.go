package k_closest_points_to_origin

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	tcases := map[string]struct {
		points [][]int
		k      int
		want   [][]int
	}{
		"k less than len(points)": {
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		"k equal to len(points)": {
			points: [][]int{{3, 3}, {-2, 4}},
			k:      2,
			want:   [][]int{{-2, 4}, {3, 3}},
		},
	}

	for tn, tc := range tcases {
		t.Run(tn, func(t *testing.T) {
			got := kClosest(tc.points, tc.k)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %+v, want %+v\n", got, tc.want)
			}
		})
	}
}
