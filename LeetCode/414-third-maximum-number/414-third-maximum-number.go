package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	unique := func(intSlice []int) []int {
		keys := make(map[int]bool)
		list := []int{}
		for _, entry := range intSlice {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		return list
	}
	uniNums := unique(nums)
	sort.Ints(uniNums)

	l := len(uniNums)
	res := uniNums[0]
	switch {
	case l > 2:
		res = uniNums[l-3]
	case l == 2:
		if res < uniNums[1] {
			res = uniNums[1]
		}
	}

	return res
}

type data struct {
	input    []int
	expected int
}

func main() {
	sols := []data{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{2}, 2},
	}

	for i, sol := range sols {
		output := thirdMax(sol.input)
		if output != sol.expected {
			fmt.Println(i)
			fmt.Println("input", sol.input)
			fmt.Println("output", output)
			fmt.Println("expect", sol.expected)
		}
	}
}
