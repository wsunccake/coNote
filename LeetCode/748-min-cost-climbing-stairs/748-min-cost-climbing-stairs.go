package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	var m int
	for i := 2; i < n; i++ {
		m = cost[i-1]
		if m > cost[i-2] {
			m = cost[i-2]
		}
		cost[i] = m + cost[i]
	}
	m = cost[n-1]
	if m > cost[n-2] {
		m = cost[n-2]
	}
	return m
}

func main() {
	var input []int
	var answer int

	input = []int{10, 15, 20}
	answer = 15
	if minCostClimbingStairs(input) != answer {
		fmt.Printf("%v != %v fail\n", input, answer)
	}

	input = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	answer = 6
	if minCostClimbingStairs(input) != answer {
		fmt.Printf("%v != %v fail\n", input, answer)
	}
}
