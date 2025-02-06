package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	tmpMap := make(map[int]int)
	max_count := 0
	max_value := 0

	for _, value := range nums {
		if _, ok := tmpMap[value]; ok {
			tmpMap[value] += 1
			if max_count < tmpMap[value] {
				max_value = value
				max_count = tmpMap[value]
			}
		} else {
			tmpMap[value] = 1
			if max_count < tmpMap[value] {
				max_value = value
				max_count = tmpMap[value]
			}
		}
	}

	return max_value
}

func main() {
	var input []int
	var answer int
	var sol int

	input = []int{3, 2, 3}
	answer = 3
	sol = majorityElement(input)
	if sol != answer {
		fmt.Printf("%d false\n", answer)
	}

	input = []int{2, 2, 1, 1, 1, 2, 2}
	answer = 2
	sol = majorityElement(input)
	if sol != answer {
		fmt.Printf("%d false\n", answer)
	}
}
