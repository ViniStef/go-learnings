package main

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, cols)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			matrix[r][c] = r * c
		}
	}

	return matrix
}
