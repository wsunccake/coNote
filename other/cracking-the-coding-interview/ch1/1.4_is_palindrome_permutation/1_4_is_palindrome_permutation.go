package main

import "fmt"

func is_palindrome_permutation(input string) bool {
	charCount := make(map[rune]int)
	totalChars := 0

	for _, char := range input {
		charCount[char]++
		totalChars++
	}

	isWordOdd := totalChars%2 == 1
	isCharOdd := false

	for _, count := range charCount {
		if count%2 != 0 {
			if isWordOdd {
				if isCharOdd {
					return false
				}
				isCharOdd = true
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	inputs := []string{"aba", "abc", "bb", "ab", "a", ""}
	outputs := []bool{true, false, true, false, true, true}
	var sol bool
	var ans bool

	for i := 0; i < len(outputs); i++ {
		sol = is_palindrome_permutation(inputs[i])
		ans = outputs[i]
		if sol != ans {
			fmt.Println(inputs[i], sol, ans)
		}
	}
}
