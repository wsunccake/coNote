package main

import (
	"fmt"
	"reflect"
)

func removeDuplicates(nums []int) int {
	current := nums[0]
	i := 1
	j := 1
	for i < len(nums) {
		if current != nums[i] {
			nums[j] = nums[i]
			j += 1
			current = nums[i]
		}
		i += 1
	}
	return j
}

type Quest struct {
	inp []int
	sol []int
	pos int
}

func checkAns(q Quest) {
	out := removeDuplicates(q.inp)
	if !reflect.DeepEqual(q.sol, q.inp[:out]) {
		fmt.Printf("%v != %v", q.inp, q.sol)
	}
}

func main() {
	var q Quest

	q = Quest{[]int{1, 1, 2}, []int{1, 2}, 2}
	checkAns(q)

	q = Quest{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5}
	checkAns(q)
}
