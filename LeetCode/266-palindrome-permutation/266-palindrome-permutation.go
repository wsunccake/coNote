package main

import "fmt"

func canPermutePalindrome(s string) bool {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	odd := 0
	for _, x := range cnt {
		odd += x & 1
	}
	return odd < 2
}

func main() {
	inputs := []string{"code", "aab", "carerac"}
	outputs := []bool{false, true, true}
	var ans bool

	for pos := 0; pos < len(outputs); pos++ {

		ans = canPermutePalindrome(inputs[pos])
		if ans != outputs[pos] {
			fmt.Println(inputs[pos], outputs[pos], ans)
		}
	}
}
