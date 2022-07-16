package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	count := 0
	n := -1
	res := false

	for i := 0; i < len(s); i++ {
		for j := n + 1; j < len(t); j++ {
			if s[i] == t[j] {
				count++
				n = j
				break
			}
		}
		if n == -1 || s[i] != t[n] {
			break
		}
	}

	if count == len(s) {
		res = true
	}

	return res
}

type data struct {
	s        string
	t        string
	expected bool
}

func main() {
	sols := []data{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"b", "c", false},
		{"", "ahbgdc", true},
		{"abc", "", false},
		{"", "", true},
		{"aaaaaa", "bbaaaa", false},
	}

	for _, sol := range sols {
		output := isSubsequence(sol.s, sol.t)
		if sol.expected != output {
			fmt.Println("error: ", sol.s, sol.t, sol.expected, output)
		}
	}

}
