package main

import "fmt"

func isUnique(input string) bool {
	seen := make(map[rune]struct{})
	for _, r := range input {
		if _, ok := seen[r]; ok {
			return false
		} else {
			seen[r] = struct{}{}
		}
	}

	return true
}

func main() {
	inputs := []string{"abcde", "hello", "apple", "kite", "padle"}
	outputs := []bool{true, false, false, true, true}
	var sol bool
	var ans bool

	for i, inp := range inputs {
		sol = isUnique(inp)
		ans = outputs[i]

		if sol != ans {
			fmt.Println(inp, sol, ans)
		}
	}
}
