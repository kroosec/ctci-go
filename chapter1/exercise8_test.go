package main

import "testing"

func TestZeroMatrix(t *testing.T) {
	testMatrix := []struct {
		input Matrix
		want  Matrix
	}{
		{Matrix{}, Matrix{}},
		{Matrix{[]int{1}}, Matrix{[]int{1}}},
		{Matrix{[]int{1, 0}}, Matrix{[]int{0, 0}}},
		{Matrix{[]int{1, 1, 1}, []int{0, 1, 1}}, Matrix{[]int{0, 1, 1}, []int{0, 0, 0}}},
		{Matrix{[]int{1, 1, 1}, []int{1, 0, 0}, []int{1, 0, 1}, []int{1, 1, 1}},
			Matrix{[]int{1, 0, 0},
				[]int{0, 0, 0},
				[]int{0, 0, 0},
				[]int{1, 0, 0}},
		},
	}
	for _, test := range testMatrix {
		got := make(Matrix, len(test.input))
		copy(got, test.input)
		ZeroMatrix(got)
		AssertInputs(t, test.input, got, test.want)
	}
}
