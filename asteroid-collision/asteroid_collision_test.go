package asteroid_collision

import (
	"slices"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	tcases := []struct {
		asteroids []int
		output    []int
	}{
		{
			asteroids: []int{5, 10, -5},
			output:    []int{5, 10},
		},
		{
			asteroids: []int{8, -8},
			output:    []int{},
		},
		{
			asteroids: []int{10, 2, -5},
			output:    []int{10},
		},
		{
			asteroids: []int{-2, -1, 1, 2},
			output:    []int{-2, -1, 1, 2},
		},
		{
			asteroids: []int{-2, -2, 1, -2},
			output:    []int{-2, -2, -2},
		},
		{
			asteroids: []int{-2, -2, 1, -1},
			output:    []int{-2, -2},
		},
		{
			asteroids: []int{-2, 1, 1, -1},
			output:    []int{-2, 1},
		},
	}

	for _, tt := range tcases {
		if output := asteroidCollision(tt.asteroids); !slices.Equal(output, tt.output) {
			t.Errorf("got %v, expected %v", output, tt.output)
		}
	}
}
