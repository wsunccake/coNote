package main

import (
	"fmt"
)

// slow
// func checkInclusion(s1 string, s2 string) bool {
// 	s1Len := len(s1)
// 	s2Len := len(s2)
// 	if s1Len > s2Len {
// 		return false
// 	}

// 	s1Map := make(map[rune]int)
// 	s2Map := make(map[rune]int)
// 	for i := 0; i < 26; i++ {
// 		s1Map[rune(97+i)] = 0
// 		s2Map[rune(97+i)] = 0
// 	}
// 	for i := 0; i < s1Len; i++ {
// 		s1Map[rune(s1[i])]++
// 		s2Map[rune(s2[i])]++
// 	}

// 	if reflect.DeepEqual(s1Map, s2Map) {
// 		return true
// 	}

// 	i := 0
// 	for j := s1Len; j < s2Len; j++ {
// 		s2Map[rune(s2[j])]++
// 		s2Map[rune(s2[i])]--
// 		if reflect.DeepEqual(s1Map, s2Map) {
// 			return true
// 		}

// 		i++
// 	}

// 	return false
// }

// fast
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
	inputs1 := []string{"ab", "ab", "abc"}
	inputs2 := []string{"eidbaooo", "eidboaoo", "ccccbbbbaaaa"}
	outputs := []bool{true, false, false}
	var ans bool

	for pos := 0; pos < len(outputs); pos++ {

		ans = checkInclusion(inputs1[pos], inputs2[pos])
		if ans != outputs[pos] {
			fmt.Println(checkInclusion(inputs1[pos], inputs2[pos]), inputs1[pos], inputs2[pos], outputs[pos], ans)
		}
	}
}
