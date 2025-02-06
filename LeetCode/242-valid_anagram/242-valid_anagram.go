package main

import (
	"fmt"
)

// // slow
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	charMap := make(map[rune]int)

// 	for _, v := range s {
// 		charMap[v] += 1
// 	}

// 	for _, v := range t {
// 		if _, ok := charMap[v]; ok {
// 			charMap[v] -= 1
// 		} else {
// 			return false
// 		}
// 	}

// 	for _, v := range charMap {
// 		if v != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

const aPos = 97

// fast
func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	list := [26]int{}

	for i := 0; i < len(s); i++ {
		list[s[i]-aPos] += 1
	}

	for i := 0; i < len(t); i++ {
		p := t[i] - aPos
		list[p] -= 1

		if list[p] < 0 {
			return false
		}
	}

	return true
}

func main() {
	inputs1 := []string{"anagram", "rat"}
	inputs2 := []string{"nagaram", "car"}
	outputs := []bool{true, false}
	var ans bool

	for pos := 0; pos < len(outputs); pos++ {

		ans = isAnagram(inputs1[pos], inputs2[pos])
		if ans != outputs[pos] {
			fmt.Println(inputs1[pos], inputs2[pos], outputs[pos], ans)
		}
	}
}
