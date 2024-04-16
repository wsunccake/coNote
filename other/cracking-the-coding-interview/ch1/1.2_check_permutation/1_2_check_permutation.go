package main

import (
	"fmt"
	"sort"
)

func isPermutation(input1, input2 string) bool {
	if len(input1) != len(input2) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range input1 {
		counts[r]++
	}
	for _, r := range input2 {
		counts[r]--
	}
	for _, val := range counts {
		if val != 0 {
			return false
		}
	}
	return true
}

// by sort
func isPermutation2(input1, input2 string) bool {
	s1 := []rune(input1)
	sort.Slice(s1, func(i int, j int) bool { return s1[i] < s1[j] })

	s2 := []rune(input2)
	sort.Slice(s2, func(i int, j int) bool { return s2[i] < s2[j] })

	return !(string(s1) != string(s2))
}

func main() {
	inputs1 := []string{"testest", "hello", "", ""}
	inputs2 := []string{"estxest", "oellh", "", "abc"}
	outputs := []bool{false, true, true, false}
	var sol bool
	var ans bool

	for i := 0; i < len(outputs); i++ {
		sol = isPermutation(inputs1[i], inputs2[i])
		ans = outputs[i]

		if sol != ans {
			fmt.Println(inputs1[i], inputs2[i], sol, ans)
		}
	}
}
