package main

import "fmt"

func truncateSentence(s string, k int) string {
	l := len(s)
	space := 0
	i := 0
	for i < l && space < k {
		if s[i] == 32 {
			space++
		}
		i++
	}

	res := s[0 : i-1]
	if space < k {
		res = s[0:i]
	}

	fmt.Println(i, l, space, k, s[0:i])
	return res
}

type data struct {
	input    string
	k        int
	expected string
}

func main() {
	sols := []data{
		{"Hello how are you Contestant", 4, "Hello how are you"},
		{"What is the solution to this problem", 4, "What is the solution"},
		{"chopper is not a tanuki", 5, "chopper is not a tanuki"},
	}

	for _, sol := range sols {
		output := truncateSentence(sol.input, sol.k)
		if sol.expected != output {
			fmt.Println(sol.input, sol.k, sol.expected, output)
		}
	}
}
