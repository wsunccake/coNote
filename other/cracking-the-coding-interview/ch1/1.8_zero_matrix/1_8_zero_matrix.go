package main

import (
	"fmt"
)

func showMatrix(input [][]int) {
	// for i1, e1 := range input {
	for _, e1 := range input {
		// for i2, e2 := range e1 {
		for _, e2 := range e1 {
			// fmt.Println(i1, i2, e2)
			fmt.Printf("%v ", e2)
		}
		fmt.Println()
	}

	for i := 0; i < len(input); i++ {
		// fmt.Println(i, input[i])
		for j := 0; j < len(input[i]); j++ {
			// fmt.Println(i, j, input[i][j])
			fmt.Printf("%v ", input[i][j])
		}
		fmt.Println()
	}
}

func zeroMatrix(input [][]int, m, n int) [][]int {
	rowLen := len(input)
	colLen := len(input[0])

	output := make([][]int, rowLen)
	for i := range output {
		output[i] = make([]int, colLen)
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if j == n {
				output[i][j] = 0
			}
			output[i][j] = input[i][j]

			if i == m {
				output[i][j] = 0
			}
		}

	}
	return output
}

func isEqual(input1, input2 [][]int) bool {
	rowLen := len(input1)
	colLen := len(input1[0])

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if input1[i][j] != input2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	inputs := [][][]int{{{1, 2, 3}, {4, 5, 6}}}
	outputs := [][][]int{{{0, 0, 0}, {0, 5, 6}}}
	var sol [][]int
	var ans [][]int

	for i := 0; i < len(outputs); i++ {
		sol = zeroMatrix(inputs[i], 0, 0)
		ans = outputs[i]

		if !isEqual(sol, ans) {
			fmt.Println(inputs[i], sol, ans)
		}
	}
}
