package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	l := len(nums)
	for i+j < l {
		if nums[i] == val {
			if nums[i+j] == val {
				j++
			} else {
				nums[i], nums[i+j] = nums[i+j], nums[i]
				i++
				j = 0

			}
		} else {
			i++
		}
	}

	return i
}

type Quest struct {
	inpNums []int
	inpVal  int
	solNums []int
	solVal  int
}

func checkAns(q Quest) {
	out := removeElement(q.inpNums, q.inpVal)
	if q.solVal != out {
		fmt.Printf("fail: sol -> %d != out -> %d", q.solVal, out)
	}
}

func main() {
	var q Quest

	q = Quest{[]int{3, 2, 2, 3}, 3, []int{2, 2}, 2}
	checkAns(q)

	q = Quest{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 4, 0, 3}, 5}
	checkAns(q)
}
