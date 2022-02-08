package main

import (
	"fmt"
	"reflect"
)

// func getConcatenation(nums []int) []int {
// 	l := len(nums)
// 	ans := make([]int, len(nums)*2)
// 	for i := 0; i < l; i++ {
// 		ans[i] = nums[i]
// 		ans[i+l] = nums[i]
// 	}
// 	return ans
// }

func getConcatenation(nums []int) []int {
	l := len(nums)
	ans := make([]int, len(nums)*2)
	for i := 0; i < 2; i++ {
		copy(ans[l*i:], nums)
	}
	return ans
}

func main() {
	i1 := []int{1, 2, 1}
	o1 := []int{1, 2, 1, 1, 2, 1}
	if !reflect.DeepEqual(getConcatenation(i1), o1) {
		fmt.Println(i1)
	}

	i2 := []int{1, 3, 2, 1}
	o2 := []int{1, 3, 2, 1, 1, 3, 2, 1}
	if !reflect.DeepEqual(getConcatenation(i2), o2) {
		fmt.Println(i2)
	}
}
