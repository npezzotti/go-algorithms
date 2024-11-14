package task_scheduler

import (
	"testing"
)

func Test_leastInterval(t *testing.T) {
	tcases := []struct {
		tasks []byte
		n     int
		res   int
	}{
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			res:   8,
		},
		{
			tasks: []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:     1,
			res:   6,
		},
		{
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     3,
			res:   10,
		},
	}

	for _, tc := range tcases {
		t.Run(string(tc.tasks), func(t *testing.T) {
			if res := leastInterval(tc.tasks, tc.n); res != tc.res {
				t.Errorf("expected %d, got %d", tc.res, res)
			}
		})
	}
}
