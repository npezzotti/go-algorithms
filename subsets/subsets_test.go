package subsets

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tcases := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{1, 2, 3},
			output: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			input:  []int{0},
			output: [][]int{{}, {0}},
		},
	}

	for _, tc := range tcases {
		output := subsets(tc.input)

		// Sort slices of slices for comparison
		sort.Slice(output, func(i, j int) bool {
			// Handle empty slices
			if len(output[i]) == 0 {
				return true // Empty slices come before non-empty ones
			}
			if len(output[j]) == 0 {
				return false
			}

			for k := 0; k < len(output[i]) && k < len(output[j]); k++ {
				if output[i][k] != output[j][k] {
					return output[i][k] < output[j][k]
				}
			}

			// If all compared elements are equal, shortest slice comes first
			return len(output[i]) < len(output[j])
		})

		if !reflect.DeepEqual(output, tc.output) {
			t.Errorf("got %+v, expected %+v", output, tc.output)
		}
	}
}
