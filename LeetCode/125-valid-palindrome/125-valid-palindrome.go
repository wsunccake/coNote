package main

import (
	"fmt"
	"strings"
)

// func isPalindrome(s string) bool {
// 	sentence := ""
// 	for _, e := range strings.ToLower(s) {
// 		if e > 96 && e < 123 {
// 			sentence = sentence + string(e)
// 		}
// 		if e > 47 && e < 58 {
// 			sentence = sentence + string(e)
// 		}
// 	}

// 	result := false
// 	sentenceLen := len(sentence)
// 	j := sentenceLen - 1
// 	for i := 0; i <= sentenceLen; i++ {
// 		if i == j-i || i > j-i {
// 			fmt.Println(sentence, i, j-i)
// 			result = true
// 			break
// 		}
// 		if sentence[i] != sentence[j-i] {
// 			break
// 		}
// 	}
// 	return result
// }

func isPalindrome(s string) bool {
	sentence := strings.ToLower(s)
	result := false
	l := len(sentence)
	i := 0
	j := l - 1
	for k := 0; k <= l; k++ {
		if i == j || i > j {
			result = true
			break
		}

		// ascii    -> char
		// 48 ~ 57  -> 0 ~ 9
		// 97 ~ 122 -> a ~ z
		if (sentence[i] < 48) || (sentence[i] > 57 && sentence[i] < 97) || (sentence[i] > 122) {
			i++
			continue
		}
		if (sentence[j] < 48) || (sentence[j] > 57 && sentence[j] < 97) || (sentence[j] > 122) {
			j--
			continue
		}

		if (sentence[i]) != sentence[j] {
			break
		}
		i++
		j--
	}
	return result
}

func main() {
	s1 := ("A man, a plan, a canal: Panama")
	// isPalindrome(s1) != true
	if !isPalindrome(s1) {
		fmt.Println(s1)
	}

	s2 := "race a car"
	// isPalindrome(s2) != false
	if isPalindrome(s2) {
		fmt.Println(s2)
	}

	s3 := " "
	if !isPalindrome(s3) {
		fmt.Println(s3, "emtpy")
	}

	s4 := "Z"
	if !isPalindrome(s4) {
		fmt.Println(s4)
	}

	s5 := "0P"
	if isPalindrome(s5) {
		fmt.Println(s5)
	}
}
