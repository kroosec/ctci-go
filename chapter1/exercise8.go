package main

// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
// column are set to 0.

type Matrix [][]int

type position struct {
	i int
	j int
}


func zeroMatrixPosition(matrix Matrix, pos position) {
	// Zero column
	for i := 0; i < len(matrix); i++ {
		matrix[i][pos.j] = 0
	}
	// Zero row
	for j := 0; j < len(matrix[pos.i]); j++ {
		matrix[pos.i][j] = 0
	}
}

// O(MxN)
func ZeroMatrix(matrix Matrix) {
	var zeroes []position
	for i, row := range matrix {
		for j, value := range row {
			if value == 0 {
				zeroes = append(zeroes, position{i, j})
			}
		}
	}
	for _, pos := range zeroes {
		zeroMatrixPosition(matrix, pos)
	}
}
