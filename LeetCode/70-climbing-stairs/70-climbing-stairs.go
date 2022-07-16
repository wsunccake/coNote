package main

import (
	"fmt"
)

// recursive
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	n0 := 1
	n1 := 1
	n2 := 1
	for i := 1; i < n; i++ {
		n2 = n1 + n0
		n0 = n1
		n1 = n2
	}
	return n2
}

// func climbStairs(n int) int {
// 	x := math.Sqrt(5)
// 	z := math.Pow((0.5+x*0.5), float64(n+1)) - math.Pow((0.5-x*0.5), float64(n+1))
// 	z = math.Floor(z/x + 0.5)
// 	return int(z)
// }

type data struct {
	x        int
	expected int
}

func main() {
	sols := []data{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{9, 55},
	}

	for _, sol := range sols {
		output := climbStairs(sol.x)
		if sol.expected != output {
			fmt.Println(sol.x, sol.expected, output)
		}
	}

}
