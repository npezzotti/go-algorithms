package cheapest_flights_within_k_stops_test

import (
	"reflect"
	"testing"
)

func Test_findCheapestPrice(t *testing.T) {
	tcases := []struct {
		flights                [][]int
		n, src, dst, k, output int
	}{
		{
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
			src:     0,
			dst:     3,
			k:       1,
			output:  700,
		},
		{
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       1,
			output:  200,
		},
		{
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       0,
			output:  500,
		},
		{
			n:       3,
			flights: [][]int{{0, 1, 200}, {1, 0, 100}, {2, 1, 500}},
			src:     0,
			dst:     2,
			k:       0,
			output:  -1,
		},
	}

	for _, tt := range tcases {
		if output := findCheapestPrice(
			tt.n, tt.flights, tt.src, tt.dst, tt.k,
		); !reflect.DeepEqual(output, tt.output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
