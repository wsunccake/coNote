package main

import (
	"fmt"
)

// func missingNumber(nums []int) int {
// 	m := map[int]int{}
// 	for _, e := range nums {
// 		m[e] = 1
// 	}
// 	r := 0
// 	isExist := true
// 	l := len(nums)
// 	for i := 0; i <= l; i++ {
// 		_, isExist = m[i]
// 		if !isExist {
// 			r = i
// 			break
// 		}
// 	}

// 	return r
// }

func missingNumber(nums []int) int {
	l := len(nums)
	sum := l * (l + 1) / 2
	for _, e := range nums {
		sum = sum - e
	}

	return sum
}

func main() {
	i := []int{3, 0, 1}
	o := 2
	if missingNumber(i) != o {
		fmt.Println(i)
	}

	i = []int{0, 1}
	o = 2
	if missingNumber(i) != o {
		fmt.Println(i)
	}

	i = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	o = 8
	if missingNumber(i) != o {
		fmt.Println(i)
	}

	i = []int{0}
	o = 1
	if missingNumber(i) != o {
		fmt.Println(i)
	}
}
