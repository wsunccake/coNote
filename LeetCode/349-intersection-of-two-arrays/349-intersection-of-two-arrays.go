package main

import (
	"fmt"
	"reflect"
	"sort"
)

// func intersection(nums1 []int, nums2 []int) []int {
// 	sorting := func(nums []int) []int {
// 		sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
// 		return nums
// 	}
// 	sorting(nums1)
// 	sorting(nums2)
// 	r := []int{}
// 	i, j := 0, 0
// 	for i < len(nums1) && j < len(nums2) {
// 		if nums1[i] > nums2[j] {
// 			i++
// 			continue
// 		}
// 		if nums1[i] < nums2[j] {
// 			j++
// 			continue
// 		}
// 		if nums1[i] == nums2[j] {
// 			if len(r) == 0 {
// 				r = append(r, nums1[i])
// 			} else {
// 				if r[len(r)-1] != nums1[i] {
// 					r = append(r, nums1[i])
// 				}
// 			}
// 			i++
// 			j++
// 			continue
// 		}
// 	}
// 	return r
// }

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	for _, num1 := range nums1 {
		m[num1] = false
	}
	for _, num2 := range nums2 {
		if _, isExist := m[num2]; isExist {
			m[num2] = true
		}
	}
	r := []int{}
	for k, v := range m {
		if v {
			r = append(r, k)
		}
	}
	return r
}

func sorted(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return nums
}

func main() {
	i := []int{1, 2, 2, 1}
	j := []int{2, 2}
	o := []int{2}
	if !reflect.DeepEqual(sorted(intersection(i, j)), sorted(o)) {
		fmt.Println(i, j)
	}

	i = []int{4, 9, 5}
	j = []int{9, 4, 9, 8, 4}
	o = []int{9, 4}
	if !reflect.DeepEqual(sorted(intersection(i, j)), sorted(o)) {
		fmt.Println(i, j)
	}

	i = []int{3, 1, 2}
	j = []int{1}
	o = []int{1}
	if !reflect.DeepEqual(sorted(intersection(i, j)), sorted(o)) {
		fmt.Println(i, j)
	}
}
