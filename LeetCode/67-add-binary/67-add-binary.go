package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func addBinary(a string, b string) string {
// 	maxLen := len(a)
// 	maxNumber := a
// 	minLen := len(b)
// 	minNumber := b
// 	if maxLen < minLen {
// 		maxNumber, minNumber = minNumber, maxNumber
// 		maxLen, minLen = minLen, maxLen
// 	}

// 	quotient := 0
// 	remainder := 0
// 	result := ""
// 	AsciiToNumber := map[uint8]int{48: 0, 49: 1}

// 	// ascii -> char
// 	// 48 	 -> 0
// 	// 49 	 -> 1
// 	for i := 0; i < minLen; i++ {
// 		switch {
// 		case maxNumber[maxLen-i-1] == 48 && maxNumber[maxLen-i-1] == 48 && quotient == 1:

// 		}
// 		remainder = (AsciiToNumber[maxNumber[maxLen-i-1]] + AsciiToNumber[minNumber[minLen-i-1]] + quotient) % 2
// 		quotient = (AsciiToNumber[maxNumber[maxLen-i-1]] + AsciiToNumber[minNumber[minLen-i-1]] + quotient) / 2
// 		result = fmt.Sprintf("%d%s", remainder, result)
// 	}

// 	for i := minLen; i < maxLen; i++ {
// 		remainder = (AsciiToNumber[maxNumber[maxLen-i-1]] + quotient) % 2
// 		quotient = (AsciiToNumber[maxNumber[maxLen-i-1]] + quotient) / 2
// 		result = fmt.Sprintf("%d%s", remainder, result)
// 	}

// 	if quotient != 0 {
// 		result = fmt.Sprintf("%d%s", quotient, result)
// 	}

// 	return result
// }

// func addBinary(a string, b string) string {
// 	maxLen := len(a)
// 	maxNumber := a
// 	minLen := len(b)
// 	minNumber := b
// 	if maxLen < minLen {
// 		maxNumber, minNumber = minNumber, maxNumber
// 		maxLen, minLen = minLen, maxLen
// 	}

// 	quotient := 0
// 	remainder := 0
// 	var results = make([]string, maxLen+1)

// 	for i := 0; i < minLen; i++ {
// 		maxVaule, _ := strconv.Atoi(maxNumber[maxLen-i-1 : maxLen-i])
// 		minVaule, _ := strconv.Atoi(minNumber[minLen-i-1 : minLen-i])
// 		remainder = (maxVaule + minVaule + quotient) % 2
// 		quotient = (maxVaule + minVaule + quotient) / 2
// 		results[maxLen-i] = strconv.Itoa(remainder)
// 	}

// 	for i := minLen; i < maxLen; i++ {
// 		maxVaule, _ := strconv.Atoi(maxNumber[maxLen-i-1 : maxLen-i])
// 		remainder = (maxVaule + quotient) % 2
// 		quotient = (maxVaule + quotient) / 2
// 		results[maxLen-i] = strconv.Itoa(remainder)
// 	}

// 	if quotient != 0 {
// 		results[0] = strconv.Itoa(quotient)
// 	}
// 	return strings.Join(results, "")
// }

func addBinary(a string, b string) string {
	s := 0
	carry := 0
	res := ""
	la := len(a) - 1
	lb := len(b) - 1
	for la >= 0 || lb >= 0 || carry != 0 {
		s = carry
		if la >= 0 {
			s += int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			s += int(b[lb] - '0')
			lb--
		}
		carry = s / 2
		res = string(s%2+'0') + res
	}
	return res
}

func main() {
	s1 := "100"
	if sol := addBinary("11", "1"); sol != s1 {
		fmt.Println("s1", s1)
	}

	s2 := "10101"
	if sol := addBinary("1010", "1011"); sol != s2 {
		fmt.Println("s2", s2)
	}

	s3 := "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"
	if sol := addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"); sol != s3 {
		fmt.Println("s3", s3)
	}

	s4 := "110001"
	if sol := addBinary("101111", "10"); sol != s4 {
		fmt.Println("s4", s4)
	}
}
