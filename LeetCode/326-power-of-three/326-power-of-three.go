package main

import (
	"fmt"
)

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	r := -1
	for n > 1 {
		r = n % 3
		n = n / 3

		if r != 0 {
			break
		}
	}

	res := false
	if r == 0 {
		res = true
	}

	return res
}

type data struct {
	input    int
	expected bool
}

func main() {
	sols := []data{
		{27, true},
		{9, true},
		{45, false},
		{0, false},
		{1, true},
		{2, false},
		{3, true},
		{-3, false},
	}

	for i, sol := range sols {
		output := isPowerOfThree(sol.input)
		if output != sol.expected {
			fmt.Println(i)
			fmt.Println("input", sol.input)
			fmt.Println("output", output)
			fmt.Println("expect", sol.expected)
		}
	}
}
