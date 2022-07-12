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

type data struct {
	input    []byte
	expected []byte
}

func main() {
	sols := []data{
		{[]byte("hi"), []byte("ih")},
		{[]byte("hiH"), []byte("Hih")},
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannnah"), []byte("hannnaH")},
		{[]byte("A man, a plan, a canal: Panama"), []byte("amanaP :lanac a ,nalp a ,nam A")},
	}

	for i, sol := range sols {
		reverseString(sol.input)
		if !reflect.DeepEqual(sol.input, sol.expected) {
			fmt.Println(i)
			fmt.Println("output", string(sol.input))
			fmt.Println("expect", string(sol.expected))
		}
	}
}
