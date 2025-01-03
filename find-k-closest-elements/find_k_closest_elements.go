package find_k_closest_elements

import (
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr) - 1
	i := sort.SearchInts(arr, x)

	l, r := i-1, i
	for k > 0 {
		k--
		if r > n || l >= 0 && (x-arr[l]) <= (arr[r]-x) {
			l--
		} else {
			r++
		}
	}

	return arr[l+1 : r]
}
