package unique_paths

import (
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tcases := []struct {
		m, n   int
		result int
	}{
		// {
		// 	m:      3,
		// 	n:      7,
		// 	result: 28,
		// },
		{
			m:      3,
			n:      2,
			result: 3,
		},
	}

	t.Run("uniquePathsRecursive", func(t *testing.T) {
		for _, tc := range tcases {
			if result := uniquePathsRecursive(tc.m, tc.n); result != tc.result {
				t.Errorf("expected %d, got %d", tc.result, result)
			}
		}
	})

	t.Run("uniquePathsDP", func(t *testing.T) {
		for _, tc := range tcases {
			if result := uniquePathsDP(tc.m, tc.n); result != tc.result {
				t.Errorf("expected %d, got %d", tc.result, result)
			}
		}
	})
}
