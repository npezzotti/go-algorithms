package course_schedule_II

import (
	"slices"
	"testing"
)

func Test_findOrder(t *testing.T) {
	tcases := []struct {
		numCourses    int
		prerequisites [][]int
		output        []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			output:        []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			output:        []int{0, 1, 2, 3},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			output:        []int{0},
		},
	}

	for _, tc := range tcases {
		if output := findOrder(tc.numCourses, tc.prerequisites); !slices.Equal(output, tc.output) {
			t.Errorf("got %v, expected %v", output, tc.output)
		}
	}
}
