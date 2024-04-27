package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	m := len(s)
	n := len(goal)
	if m > n {
		return false
	}
	if strings.Contains(s+s, goal) {
		return true
	}
	return false
}

// func rotateString(s string, goal string) bool {
// 	m := len(s)
// 	n := len(goal)
// 	if m > n {
// 		return false
// 	}
// 	for i := 0; i < m; i++ {
// 		if s[i:]+s[:i] == goal {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {
	inputs1 := []string{"abcde", "abcde", "aa"}
	inputs2 := []string{"cdeab", "abced", "a"}
	outputs := []bool{true, false, false}

	for pos := 0; pos < len(outputs); pos++ {
		ans := rotateString(inputs1[pos], inputs2[pos])
		if ans != outputs[pos] {
			fmt.Println(inputs1[pos], inputs2[pos], ans, outputs[pos])
		}
	}

}
