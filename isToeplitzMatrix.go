func isToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		if !valid(matrix, i, 0) {
			return false
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if !valid(matrix, 0, i) {
			return false
		}
	}
	return true
}

func valid(matrix [][]int, i int, j int) bool {
	for i < len(matrix)-1 && j < len(matrix[0])-1 {
		if matrix[i][j] != matrix[i+1][j+1] {
			return false
		}
		i++
		j++
	}
	return true
}