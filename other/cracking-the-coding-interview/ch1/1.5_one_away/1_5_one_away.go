package main

import "fmt"

func maxString(input1, input2 string) (string, string) {
	s1, s2 := input1, input2
	if len(input1) < len(input2) {
		s1, s2 = input2, input1
	}
	return s1, s2
}

// input1 ,input2 same length
func isDifferent(input1, input2 string) int {
	diffCount := 0
	for i := 0; i < len(input1); i++ {
		if input1[i] != input2[i] {
			diffCount++
		}
	}
	return diffCount
}

func isDifferentOneChar(input1, input2 string) int {
	// fmt.Println(input1, input2)
	diffCount := 0
	i := 0
	len1 := len(input1)
	len2 := len(input2)
	for j := 0; j < len2; j++ {
		for ; i < len1; i++ {
			// fmt.Println("i j ", i, j, input1[i], input2[j], diffCount)
			if input1[i] == input2[j] {
				break
			} else {
				diffCount++
			}
		}
		i++

	}
	if i < len1 {
		diffCount += len1 - i
	}
	// fmt.Println(diffCount)

	return diffCount
}

func oneAway(input1, input2 string) bool {
	s1, s2 := maxString(input1, input2)

	len1 := len(s1)
	len2 := len(s2)

	if !(len1 == len2 || len1-1 == len2) {
		return false
	}

	if len1 == len2 {
		if isDifferent(s1, s2) > 1 {
			return false
		}
	} else {
		if isDifferentOneChar(s1, s2) > 1 {
			return false
		}
	}

	return true
}

func main() {
	inputs1 := []string{"pale", "pales", "pale", "pale",
		"abcd", "abcd", "abcd", "abcd", "abcd", "abcde", "abcdf",
		" ", "", "",
	}
	inputs2 := []string{"ple", "pale", "bale", "bake",
		"abcd", "abcc", "accc", "abcde", "abcdef", "abcd", "abcd",
		"", " ", "",
	}
	outputs := []bool{true, true, true, false,
		true, true, false, true, false, true,
		true, true, true,
	}
	var sol bool
	var ans bool

	for i := 0; i < len(outputs); i++ {
		sol = oneAway(inputs1[i], inputs2[i])
		ans = outputs[i]
		if sol != ans {
			fmt.Println(inputs1[i], inputs2[i], sol, ans)
		}
		// fmt.Println(inputs1[i], inputs2[i], sol, ans)
	}
}
