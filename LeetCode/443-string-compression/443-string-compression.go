package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	count := 1
	tmpStr := string(chars[0])

	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1] {
			count += 1
		} else {
			if count != 1 {
				tmpStr += strconv.Itoa(count)
			}
			tmpStr += string(chars[i+1])
			count = 1
		}
	}

	if count != 1 {
		// tmpStr += fmt.Sprint(count)
		tmpStr += strconv.Itoa(count)
	}

	fmt.Println(tmpStr)
	for i := 0; i < (len(tmpStr)); i++ {
		chars[i] = tmpStr[i]
	}

	return len(tmpStr)
}

func compress1(chars []byte) int {
	char_slice := []byte{chars[0]}
	num_slice := []int{1}

	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1] {
			num_slice[len(char_slice)-1] += 1
		} else {
			char_slice = append(char_slice, chars[i+1])
			num_slice = append(num_slice, 1)
		}
	}

	tmp_str := ""
	for i := 0; i < (len(char_slice)); i++ {
		tmp_str = fmt.Sprintf("%s%c", tmp_str, char_slice[i])
		if num_slice[i] != 1 {
			tmp_str = fmt.Sprintf("%s%d", tmp_str, num_slice[i])
		}
	}

	for pos, c := range tmp_str {
		chars[pos] = byte(c)
	}

	return len(tmp_str)
}

func main() {
	inputs := [][]byte{
		{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		{'a'},
		{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		{'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'a', 'b'},
	}
	outputs := [][]byte{
		{'a', '2', 'b', '2', 'c', '3'},
		{'a'},
		{'a', 'b', '1', '2'},
		{'b', '1', '1', 'a', 'b'},
	}
	nums := []int{6, 1, 4, 5}
	var ans int

	for pos := 0; pos < len(nums); pos++ {

		ans = compress(inputs[pos])
		// if inputs[pos][0:ans] != outputs[pos] {
		fmt.Println(inputs[pos], outputs[pos], ans, nums[pos])
		// }
	}
}
