package main

import "math"

// Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
// bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

// Matrix is represented as by a N^2 int array.

func FlipMatrixInPlace(matrix []int) {
	if len(matrix) <= 1 {
		return
	}
	N := int(math.Sqrt(float64(len(matrix))))
	if N*N != len(matrix) {
		// Erroneous matrix array size
		return
	}
	for layer := 0; layer < N/2; layer++ {
		first := layer
		last := N - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first*N+i]
			matrix[first*N+i] = matrix[(last-offset)*N+first]
			matrix[(last-offset)*N+first] = matrix[last*N+last-offset]
			matrix[last*N+last-offset] = matrix[i*N+last]
			matrix[i*N+last] = top
		}
	}
}

func FlipMatrix(matrix []int) []int {
	newMatrix := make([]int, len(matrix))
	copy(newMatrix, matrix)
	if len(matrix) <= 1 {
		return newMatrix
	}
	N := int(math.Sqrt(float64(len(matrix))))
	if N*N != len(matrix) {
		// Erroneous matrix array size
		return nil
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			newMatrix[N-1-i+j*N] = matrix[j+i*N]
		}
	}
	return newMatrix
}
