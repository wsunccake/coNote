package main

import "fmt"

func largestOddNumber(num string) string {
	l := len(num)
	i := l
	for i > 0 {
		i--
		r := num[i] % 2
		if r == 1 {
			break
		}
	}

	res := ""
	if i > 0 {
		res = num[0 : i+1]
	}
	if i == 0 && (num[0]%2 == 1) {
		res = num[0:1]
	}

	return res
}

type data struct {
	input    string
	expected string
}

func main() {
	sols := []data{
		{"52", "5"},
		{"4206", ""},
		{"35427", "35427"},
	}

	for _, sol := range sols {
		output := largestOddNumber(sol.input)
		if sol.expected != output {
			fmt.Println(sol.input, sol.expected, output)
		}
	}
}
