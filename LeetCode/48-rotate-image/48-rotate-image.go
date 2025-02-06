package main

import (
	"fmt"
	"strconv"
)

func rotate(matrix [][]int) {
	rawMap := make(map[string]int)

	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rawMap[strconv.Itoa(i)+"_"+strconv.Itoa(j)] = matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = rawMap[strconv.Itoa(n-j-1)+"_"+strconv.Itoa(i)]
		}
	}
}

func main() {
	inputs := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}
	outputs := [][][]int{
		{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	}

	for pos := 0; pos < len(inputs); pos++ {

		fmt.Println(inputs[pos], outputs[pos])
		rotate(inputs[pos])
		// if inputs[pos][0:ans] != outputs[pos] {
		fmt.Println(inputs[pos], outputs[pos])
		// }
	}
}
