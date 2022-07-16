package main

import "fmt"

func fib(n int) int {
	f := []int{0, 1}
	for i := 2; i <= n; i++ {
		f = append(f, f[i-2]+f[i-1])
	}
	return f[n]
}

type Quest struct {
	inp int
	sol int
}

func checkAns(q Quest) {
	out := fib(q.inp)
	if q.sol != out {
		fmt.Printf("%v, fail: sol -> %d != out -> %d", q.inp, q.sol, out)
	}
}

func main() {
	var q Quest

	q = Quest{2, 1}
	checkAns(q)

	q = Quest{3, 2}
	checkAns(q)

	q = Quest{4, 3}
	checkAns(q)
}
