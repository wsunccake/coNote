package main

import "fmt"

func convertToTitle(columnNumber int) string {
	number_to_char_map := make(map[int]string)
	for i := 65; i < 90; i++ {
		number_to_char_map[i-64] = string(i)
	}
	number_to_char_map[0] = "Z"

	s := ""
	for columnNumber > 0 {
		remainder := columnNumber % 26
		s = fmt.Sprintf("%s%s", number_to_char_map[remainder], s)
		columnNumber = (columnNumber - 1) / 26
	}

	return s
}

type Quest struct {
	inp int
	sol string
}

func checkAns(q Quest) bool {
	out := convertToTitle(q.inp)
	res := (out == q.sol)
	if !res {
		fmt.Printf("%v -> %v != %v\n", q.inp, q.sol, out)
	}
	return res
}

func main() {
	var q *Quest

	q = &Quest{1, "A"}
	checkAns(*q)

	q = &Quest{26, "Z"}
	checkAns(*q)

	q = &Quest{27, "AA"}
	checkAns(*q)

	q = &Quest{28, "AB"}
	checkAns(*q)

	q = &Quest{52, "AZ"}
	checkAns(*q)
}
