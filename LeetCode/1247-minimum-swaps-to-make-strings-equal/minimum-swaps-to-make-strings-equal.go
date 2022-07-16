package main

import (
	"fmt"
)

func minimumSwap(s1 string, s2 string) int {
	xy := 0
	yx := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == 'x' && s2[i] == 'y' {
			xy++
		}

		if s1[i] == 'y' && s2[i] == 'x' {
			yx++
		}
	}

	if (xy+yx)%2 != 0 {
		return -1
	}
	return (xy+1)/2 + (yx+1)/2
}

type data struct {
	input1   string
	input2   string
	expected int
}

func main() {
	sols := []data{
		{"xx", "yy", 1},
		{"xy", "yx", 2},
		{"xy", "xx", -1},
		{"xxyyxyxyxx", "xyyxyxxxyx", 4},
	}

	for _, sol := range sols {
		output := minimumSwap(sol.input1, sol.input2)
		if sol.expected != output {
			fmt.Println(sol.input1, sol.input2, output, sol.expected)
		}
	}

}
