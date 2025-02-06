package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	type p struct{ x, y int }
	zeros := []p{}
	l1 := len(matrix)
	l2 := len(matrix[0])

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, p{i, j})
			}
		}
	}

	for _, e := range zeros {
		for i := 0; i < l1; i++ {
			matrix[i][e.y] = 0
		}
		for j := 0; j < l2; j++ {
			matrix[e.x][j] = 0
		}
	}
}

func twoDimensionEqual(d1 [][]int, d2 [][]int) bool {
	res := true
	for i := 0; i < len(d1); i++ {
		for j := 0; j < len(d1[0]); j++ {
			if d1[i][j] != d2[i][j] {
				res = false
				break
			}
		}
		if !res {
			break
		}
	}
	return res
}

func main() {
	inputs := [][][]int{
		{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
	}
	outputs := [][][]int{
		{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
	}

	for pos := 0; pos < len(outputs); pos++ {
		setZeroes(inputs[pos])
		if !twoDimensionEqual(inputs[pos], outputs[pos]) {
			fmt.Println(inputs[pos], outputs[pos])
		}
	}
}
