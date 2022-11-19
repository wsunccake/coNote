package main

import "fmt"

func bestHand(ranks []int, suits []byte) string {
	suitsMap := make(map[byte]int)
	var ok bool
	for _, v := range suits {
		_, ok = suitsMap[v]
		if ok {
			suitsMap[v] += 1
		} else {
			suitsMap[v] = 1
		}
	}
	if len(suitsMap) == 1 {
		return "Flush"
	}

	ranksMap := make(map[int]int)
	max := 0
	for _, v := range ranks {
		_, ok = ranksMap[v]
		if ok {
			ranksMap[v] += 1
		} else {
			ranksMap[v] = 1
		}
		if max < ranksMap[v] {
			max = ranksMap[v]
		}
	}
	resultMap := map[int]string{
		4: "Three of a Kind",
		3: "Three of a Kind",
		2: "Pair",
	}

	r := "High Card"
	var s string
	s, ok = resultMap[max]
	if ok {
		r = s
	}

	return r
}

func main() {
	var ranks []int
	var suits []byte
	var answer string
	var sol string

	ranks = []int{13, 2, 3, 1, 9}
	suits = []byte{'a', 'a', 'a', 'a', 'a'}
	answer = "Flush"
	sol = bestHand(ranks, suits)
	if sol != answer {
		fmt.Printf("%s false\n", answer)
	}

	ranks = []int{4, 4, 2, 4, 4}
	suits = []byte{'d', 'a', 'a', 'b', 'c'}
	answer = "Three of a Kind"
	sol = bestHand(ranks, suits)
	if sol != answer {
		fmt.Printf("%s false\n", answer)
	}

	ranks = []int{10, 10, 2, 12, 9}
	suits = []byte{'a', 'b', 'c', 'a', 'd'}
	answer = "Pair"
	sol = bestHand(ranks, suits)
	if sol != answer {
		fmt.Printf("%s false\n", answer)
	}

	ranks = []int{2, 10, 7, 10, 7}
	suits = []byte{'a', 'b', 'a', 'd', 'b'}
	answer = "Pair"
	sol = bestHand(ranks, suits)
	if sol != answer {
		fmt.Printf("%s false\n", answer)
	}
}
