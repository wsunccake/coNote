package main

import (
	"fmt"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	for j := 1; j < n+1; j++ {
		dp[0][j] = j
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(Min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				// for go 1.21
				// dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func checkInclusion(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len > s2Len {
		return false
	}

	s1Arr := [26]int{}
	s2Arr := [26]int{}
	for i := 0; i < s1Len; i++ {
		s1Arr[rune(s1[i])-rune('a')]++
		s2Arr[rune(s2[i])-rune('a')]++
	}

	if s1Arr == s2Arr {
		return true
	}

	i := 0
	for j := s1Len; j < s2Len; j++ {
		s2Arr[rune(s2[j])-rune('a')]++
		s2Arr[rune(s2[i])-rune('a')]--

		if s1Arr == s2Arr {
			return true
		}
		i++
	}

	return false
}

func main() {
	// minDistance("ab", "a")
	inputs1 := []string{"horse", "intention"}
	inputs2 := []string{"ros", "execution"}
	outputs := []int{3, 5}
	var ans int

	for pos := 0; pos < len(outputs); pos++ {

		ans = minDistance(inputs1[pos], inputs2[pos])
		if ans != outputs[pos] {
			fmt.Println(checkInclusion(inputs1[pos], inputs2[pos]), inputs1[pos], inputs2[pos], outputs[pos], ans)
		}
	}
}
