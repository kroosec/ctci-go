package main

import "testing"

func TestFlipMatrix(t *testing.T) {
	testInputs := []struct {
		matrix []int
		want   []int
	}{
		{[]int{}, []int{}},                                                   // N = 0
		{[]int{1}, []int{1}},                                                 // N = 1
		{[]int{1, 2, 3, 4}, []int{3, 1, 4, 2}},                               // N = 2
		{[]int{1, 2, 3, 4}, []int{3, 1, 4, 2}},                               // N = 2
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{7, 4, 1, 8, 5, 2, 9, 6, 3}}, // N = 3
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			[]int{13, 9, 5, 1, 14, 10, 6, 2, 15, 11, 7, 3, 16, 12, 8, 4}}, // N = 4
	}
	for _, input := range testInputs {
		got := FlipMatrix(input.matrix)
		AssertInputs(t, input.matrix, got, input.want)
		got = make([]int, len(input.matrix))
		copy(got, input.matrix)
		FlipMatrixInPlace(got)
		AssertInputs(t, input.matrix, got, input.want)
	}
}
