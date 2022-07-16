package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	patDict := map[string]bool{}
	wordPatternDict := map[string]string{}
	res := true

	if len(strings.Split(s, " ")) != len(pattern) {
		return false
	}

	for i, word := range strings.Split(s, " ") {
		p := pattern[i : i+1]
		_, patExist := patDict[p]
		val, wordPatternExist := wordPatternDict[word]
		if wordPatternExist {
			if val != p {
				res = false
				break
			}
		} else {
			if patExist {
				res = false
				break
			} else {
				wordPatternDict[word] = p
			}
		}

		patDict[p] = true
	}

	return res
}

// func wordPattern(pattern string, s string) bool {
// 	hash := make(map[byte]string)
// 	hash2 := make(map[string]byte)
// 	result := strings.Split(s, " ")
// 	if len(result) != len(pattern) {
// 		return false
// 	}
// 	for i := 0; i < len(result); i++ {
// 		value, ok := hash[pattern[i]]
// 		value2, ok2 := hash2[result[i]]
// 		if !ok && !ok2 {
// 			hash[pattern[i]] = result[i]
// 			hash2[result[i]] = pattern[i]
// 		} else {
// 			if value == result[i] && value2 == pattern[i] {
// 				continue
// 			} else {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

func main() {
	i1 := "abba"
	i2 := "dog cat cat dog"
	o := wordPattern(i1, i2)
	if !o {
		fmt.Println(i1, i2, o)
	}

	i1 = "abba"
	i2 = "dog cat cat fish"
	o = wordPattern(i1, i2)
	if o {
		fmt.Println(i1, i2, o)
	}

	i1 = "aaaa"
	i2 = "dog cat cat dog"
	o = wordPattern(i1, i2)
	if o {
		fmt.Println(i1, i2, o)
	}

	i1 = "abba"
	i2 = "dog dog dog dog"
	o = wordPattern(i1, i2)
	if o {
		fmt.Println(i1, i2, o)
	}

	i1 = "aaa"
	i2 = "aa aa aa aa"
	o = wordPattern(i1, i2)
	if o {
		fmt.Println(i1, i2, o)
	}

	i1 = "jquery"
	i2 = "jquery"
	o = wordPattern(i1, i2)
	if o {
		fmt.Println(i1, i2, o)
	}
}
