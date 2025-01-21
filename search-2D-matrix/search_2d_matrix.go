package search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	b, t := 0, len(matrix)-1
	for b < t {
		mid := ((b + t) / 2) + 1

		if target == matrix[mid][0] {
			return true
		} else if target > matrix[mid][0] {
			b = mid
		} else {
			t = mid - 1
		}
	}

	l, r := 0, len(matrix[b])-1
	for l <= r {
		mid := (l + r) / 2

		if matrix[b][mid] == target {
			return true
		} else if target > matrix[b][mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
