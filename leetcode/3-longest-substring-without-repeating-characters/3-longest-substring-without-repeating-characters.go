package main

import (
	"fmt"
)

// func lengthOfLongestSubstring(s string) int {
// 	maxStr := ""
// 	for i := 0; i < len(s); i++ {
// 		tmpStr := ""
// 		tmpM := map[string]bool{}
// 		for j := i; j < len(s); j++ {
// 			subStr := s[j : j+1]
// 			if _, isExist := tmpM[subStr]; isExist {
// 				break
// 			} else {
// 				tmpM[subStr] = true
// 				tmpStr = tmpStr + subStr
// 			}
// 		}
// 		if len(maxStr) < len(tmpStr) {
// 			maxStr = tmpStr
// 		}
// 	}
// 	return len(maxStr)
// }

func lengthOfLongestSubstring(s string) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	maxLen := 0
	tmpM := map[uint8]int{}
	begin, end := 0, 0
	for end < len(s) {
		if _, isExist := tmpM[s[end]]; isExist {
			delete(tmpM, s[begin])
			begin++
		} else {
			tmpM[s[end]] = end
			end++
			maxLen = max(maxLen, len(tmpM))
		}
	}
	return maxLen
}

func main() {
	i := "abcabcbb"
	o := 3
	if lengthOfLongestSubstring(i) != o {
		fmt.Println(i, "abc")
	}

	i = "bbbbb"
	o = 1
	if lengthOfLongestSubstring(i) != o {
		fmt.Println(i, "b")
	}

	i = "pwwkew"
	o = 3
	if lengthOfLongestSubstring(i) != o {
		fmt.Println(i, "wke")
	}

	i = ""
	o = 0
	if lengthOfLongestSubstring(i) != o {
		fmt.Println(i, "")
	}

	i = "aab"
	o = 2
	if lengthOfLongestSubstring(i) != o {
		fmt.Println(i, "ab")
	}

	i = "dvdf"
	o = 3
	if lengthOfLongestSubstring(i) != o {
		fmt.Println(i, "vdf")
	}
}
