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
    sol1 := twoSum([]int{1 ,2 ,3}, 6)
    if sol1[0] != 0  { fmt.Println(false) }
    if sol1[1] != 0  { fmt.Println(false) }

    sol2 := twoSum([]int{3 ,2 ,4}, 6)
    if sol2[0] != 1  { fmt.Println(false) }
    if sol2[1] != 2  { fmt.Println(false) }

    sol3 := twoSum([]int{3 ,3}, 6)
    if sol3[0] != 0  { fmt.Println(false) }
    if sol3[1] != 1  { fmt.Println(false) }
}

