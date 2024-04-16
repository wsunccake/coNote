package main

import (
	"fmt"
	"strconv"
)

func stringCompress(input string) string {
	compressed := ""
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			compressed += string(input[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	compressed += string(input[len(input)-1]) + strconv.Itoa(count)

	return compressed
}

func stringCompression(input string) string {
	result := input
	if len(input) > len(stringCompress(input)) {
		result = stringCompress(input)
	}
	return result
}
func main() {

	inputs := []string{"aabcccccaaa", "abc"}
	outputs := []string{"a2b1c5a3", "abc"}
	var sol string
	var ans string

	for i := 0; i < len(outputs); i++ {
		sol = stringCompression(inputs[i])
		ans = outputs[i]

		if sol != ans {
			fmt.Println(inputs[i], sol, ans)
		}
	}
}
