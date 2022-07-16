package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen == 0 {
		return 0
	}

	match := -1
	for i := 0; i <= haystackLen-needleLen; i++ {
		if haystack[i] == needle[0] && haystack[i+needleLen-1] == needle[needleLen-1] {
			for j := 0; j < needleLen; j++ {
				if haystack[i+j] != needle[j] {
					match = -1
					break
				}
				match = i
			}
			if match != -1 {
				break
			}
		}
	}
	return match
}

func main() {
	i := "hello"
	j := "ll"
	if k := strStr(i, j); k != 2 {
		fmt.Println(i, j, k)
	}

	i = "aaaaa"
	j = "bba"
	if k := strStr(i, j); k != -1 {
		fmt.Println(i, j, k)
	}

	i = ""
	j = ""
	if k := strStr(i, j); k != 0 {
		fmt.Println(i, j, k)
	}

	i = ""
	j = "a"
	if k := strStr(i, j); k != -1 {
		fmt.Println(i, j, k)
	}

	i = "a"
	j = "a"
	if k := strStr(i, j); k != 0 {
		fmt.Println(i, j, k)
	}

	i = "aaa"
	j = "a"
	if k := strStr(i, j); k != 0 {
		fmt.Println(i, j, k)
	}

	i = "mississippi"
	j = "sipp"
	if k := strStr(i, j); k != 6 {
		fmt.Println(i, j, k)
	}
}
