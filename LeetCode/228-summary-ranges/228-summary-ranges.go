package main

import "fmt"

func summaryRanges(nums []int) []string {
	l := len(nums)
	results := []string{}
	if l == 0 {
		return results
	}

	beginNum := nums[0]
	endNum := nums[0]
	output := func(x, y int) string {
		r := fmt.Sprint(x)
		if x != y {
			r = fmt.Sprintf("%d->%d", x, y)
		}
		return r
	}

	for i := 0; i < l; i++ {
		if endNum+1 < nums[i] {
			results = append(results, output(beginNum, endNum))
			beginNum = nums[i]
		}
		endNum = nums[i]
	}
	results = append(results, output(beginNum, endNum))

	return results
}

func checkAnswer(sol []string, ans []string) {
	if len(sol) != len(ans) {
		fmt.Printf("%v != %v : ", sol, ans)
		fmt.Println(false)
	}
	for i := 0; i < len(sol); i++ {
		if sol[i] != ans[i] {
			fmt.Printf("%s != %s: ", sol[i], ans[i])
			fmt.Println(false)
		}
	}
}

func main() {
	sol1 := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	ans1 := []string{"0->2", "4->5", "7"}
	checkAnswer(sol1, ans1)

	sol2 := summaryRanges([]int{0, 2, 3, 3, 4, 6, 8, 9})
	ans2 := []string{"0", "2->4", "6", "8->9"}
	checkAnswer(sol2, ans2)

	sol3 := summaryRanges([]int{})
	ans3 := []string{}
	checkAnswer(sol3, ans3)

	sol4 := summaryRanges([]int{-1, -1})
	ans4 := []string{"-1"}
	checkAnswer(sol4, ans4)

	sol5 := summaryRanges([]int{-1})
	ans5 := []string{"-1"}
	checkAnswer(sol5, ans5)

}
