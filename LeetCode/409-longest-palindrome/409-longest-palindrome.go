package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	stringMap := make(map[int32]int)
	var ok bool

	for _, e := range s {

		_, ok = stringMap[e]
		if ok {
			stringMap[e] += 1
		} else {
			stringMap[e] = 1
		}
	}

	r := 0
	q := 0
	for _, v := range stringMap {
		r += int(v / 2)
		if v%2 == 1 {
			q = 1
		}

	}

	return r*2 + q
}

func main() {
	var input string
	var answer int

	input = "abccccdd"
	answer = 7
	if longestPalindrome(input) != answer {
		fmt.Printf("%v != %v fail\n", input, answer)
	}

	input = "a"
	answer = 1
	if longestPalindrome(input) != answer {
		fmt.Printf("%v != %v fail\n", input, answer)
	}
}
