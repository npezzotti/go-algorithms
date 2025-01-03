package find_k_closest_elements

func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k

	for l < r {
		mid := (l + r) >> 1

		if x-arr[mid] <= arr[mid+k]-x {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return arr[l : l+k]
}
