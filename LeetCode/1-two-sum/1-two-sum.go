package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    result := []int{0, 0}
    tmpMap := make(map[int]int)

    for index, element := range(nums) {
        var goal = target - element
        if goal_index, ok := tmpMap[goal]; ok {
            return []int{goal_index, index}
        }

        tmpMap[element] = index
    }

    return result
}


func main() {
    var input1 []int
    var input2 int
    var answer1 int
    var answer2 int
    var sol []int

    input1 = []int{1, 2, 3}
    input2 = 6
    answer1 = 0
    answer2 = 0
    sol = twoSum(input1, input2)
    if sol[0] != answer1  { fmt.Printf("%d false\n", answer1) }
    if sol[1] != answer2  { fmt.Printf("%d false\n", answer2) }

    input1 = []int{3, 2, 4}
    input2 = 6
    answer1 = 1
    answer2 = 2
    sol = twoSum(input1, input2)
    if sol[0] != answer1  { fmt.Printf("%d false\n", answer1) }
    if sol[1] != answer2  { fmt.Printf("%d false\n", answer2) }

    input1 = []int{3, 3}
    input2 = 6
    answer1 = 0
    answer2 = 1
    sol = twoSum(input1, input2)
    if sol[0] != answer1  { fmt.Printf("%d false\n", answer1) }
    if sol[1] != answer2  { fmt.Printf("%d false\n", answer2) }
}

