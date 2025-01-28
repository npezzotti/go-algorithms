package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	var i0, j0 bool
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				if i == 0 {
					i0 = true
				} else {
					matrix[i][0] = 0
				}

				if j == 0 {
					j0 = true
				} else {
					matrix[0][j] = 0
				}
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := range matrix {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			matrix[i] = make([]int, len(matrix[0]))
		}
	}


	if i0 {
		matrix[0] = make([]int, len(matrix[0]))
	}
	
	if j0 {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
