package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	tcases := []struct {
		numCourses    int
		prerequisites [][]int
		canFinish     bool
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			canFinish:     true,
		},
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			canFinish:     false,
		},
		{
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {2, 0}},
			canFinish:     true,
		},
	}

	for _, tc := range tcases {
		got := canFinish(tc.numCourses, tc.prerequisites)
		if got != tc.canFinish {
			t.Fatalf("got %v, want %v", got, tc.canFinish)
		}
	}
}
