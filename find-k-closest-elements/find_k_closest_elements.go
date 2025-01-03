package find_k_closest_elements

import "math"

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k

	for l < r {
		mid := (l + r) >> 1

		if math.Abs(float64(x-arr[mid])) <= math.Abs(float64(x-arr[mid+k])) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return arr[l : l+k]
}
