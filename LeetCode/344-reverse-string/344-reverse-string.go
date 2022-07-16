package main

import (
	"fmt"
	"reflect"
)

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < (l+1)/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

type Quest struct {
	inp []byte
	sol []byte
}

func checkAns(q Quest) {
	reverseString(q.inp)
	if !reflect.DeepEqual(q.sol, q.inp) {
		fmt.Printf("%v != %v", q.inp, q.sol)
	}
}

func main() {
	var q Quest

	q = Quest{[]byte("hi"), []byte("ih")}
	checkAns(q)

	q = Quest{[]byte("hiH"), []byte("Hih")}
	checkAns(q)

	q = Quest{[]byte("hello"), []byte("olleh")}
	checkAns(q)

	q = Quest{
		[]byte("A man, a plan, a canal: Panama"),
		[]byte("amanaP :lanac a ,nalp a ,nam A"),
	}
	checkAns(q)
}
