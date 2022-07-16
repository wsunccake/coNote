package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	c := head

	n := head.Next

	for n != nil {
		if c.Val == n.Val {
			c.Next = n.Next
		} else {
			c = n
		}
		n = n.Next
	}

	return head
}

func createListNode(inputs []int) *ListNode {
	l := len(inputs)
	if l == 0 {
		return nil
	}
	head := &ListNode{inputs[0], nil}
	c := head
	for i := 1; i < l; i++ {
		c.Next = &ListNode{inputs[i], nil}
		c = c.Next
	}

	return head
}

func traverseListNode(head ListNode) string {
	s := "node:"
	c := &head
	for c != nil {
		s = fmt.Sprintf("%s %d", s, c.Val)
		c = c.Next
	}
	return s
}

type Quest struct {
	inp *ListNode
	sol *ListNode
}

func checkAns(q Quest) {
	sol := "nil"
	if q.sol != nil {
		sol = traverseListNode(*q.sol)
	}

	ans := deleteDuplicates(q.inp)
	out := "nil"
	if ans != nil {
		out = traverseListNode(*ans)
	}

	if sol != out {
		fmt.Printf("fail: sol -> %s != out -> %s", sol, out)
	}
}

func main() {
	q1 := Quest{
		createListNode([]int{1, 1, 2}),
		createListNode([]int{1, 2}),
	}
	checkAns(q1)

	q2 := Quest{
		createListNode([]int{1, 1, 2, 3, 3}),
		createListNode([]int{1, 2, 3}),
	}
	checkAns(q2)

	q3 := Quest{
		createListNode([]int{1}),
		createListNode([]int{1}),
	}
	checkAns(q3)

	q4 := Quest{
		createListNode([]int{}),
		createListNode([]int{}),
	}
	checkAns(q4)
}
