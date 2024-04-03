package main

import "fmt"

func URLify(input string) string {
	r := []rune{}

	for _, e := range input {

		if e == 32 {
			r = append(r, '%')
			r = append(r, '2')
			r = append(r, '0')
		} else {
			r = append(r, e)
		}

	}
	return string(r)
}

func main() {
	inputs := []string{"Mr 3ohn Smith"}
	outputs := []string{"Mr%203ohn%20Smith"}
	var sol string
	var ans string

	for i := 0; i < len(outputs); i++ {
		sol = URLify(inputs[i])
		ans = outputs[i]
		if sol != ans {
			fmt.Println(inputs[i], sol, ans)
		}
	}
}
