package main

import "fmt"

func containsDuplicate(nums []int) bool {
	isDuplicate := false
	numMap := make(map[int]int)

	for _, val := range nums {
		if _, ok := numMap[val]; ok {
			isDuplicate = true
			break
		} else {
			numMap[val] = 1
		}
	}

	return isDuplicate
}

func main() {
	inputs := [][]int{
		[]int{1, 2, 3, 1},
		[]int{1, 2, 3, 4},
		[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	outputs := []bool{true, false, true}
	var ans bool

	for pos := 0; pos < len(outputs); pos++ {

		ans = containsDuplicate(inputs[pos])
		if ans != outputs[pos] {
			fmt.Println(inputs[pos], outputs[pos], ans)
		}
	}
}
