package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	res := 0
	for i := 0; i < l/2; i++ {
		res += nums[2*i]
	}
	return res
}

type data struct {
	input    []int
	expected int
}

func main() {
	sols := []data{
		{[]int{1, 4}, 1},
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
	}

	for i, sol := range sols {
		output := arrayPairSum(sol.input)
		if output != sol.expected {
			fmt.Println(i)
			fmt.Println("input", sol.input)
			fmt.Println("output", output)
			fmt.Println("expect", sol.expected)
		}
	}
}
