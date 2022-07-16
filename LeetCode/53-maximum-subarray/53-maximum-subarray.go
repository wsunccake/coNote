package main

import "fmt"

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	sub_sum := 0
	for i := 0; i < len(nums); i++ {
		if sub_sum > 0 {
			sub_sum += nums[i]
		} else {
			sub_sum = nums[i]
		}

		if sub_sum > max_sum {
			max_sum = sub_sum
		}
	}
	return max_sum
}

type Quest struct {
	inp []int
	sol int
}

func checkAns(q Quest) {
	out := maxSubArray(q.inp)
	if q.sol != out {
		fmt.Printf("%v, fail: sol -> %d != out -> %d", q.inp, q.sol, out)
	}
}

func main() {
	var q Quest

	q = Quest{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}
	checkAns(q)

	q = Quest{[]int{1}, 1}
	checkAns(q)

	q = Quest{[]int{5, 4, -1, 7, 8}, 23}
	checkAns(q)

	q = Quest{[]int{-10, -3, -1, -4}, -1}
	checkAns(q)
}
