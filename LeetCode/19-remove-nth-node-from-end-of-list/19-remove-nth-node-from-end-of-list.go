package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}

	head := &ListNode{Val: l[0]}
	cur := head
	for i := 1; i < len(l); i++ {
		cur.Next = &ListNode{Val: l[i]}
		cur = cur.Next
	}

	return head
}

func showList(head *ListNode) string {
	r := ""
	if head == nil {
		return r
	}

	cur := head
	for cur.Next != nil {
		r += fmt.Sprintf("%d->", cur.Val)
		cur = cur.Next
	}
	r += fmt.Sprintf("%d", cur.Val)
	return r
}

//	func showList(head *ListNode) string {
//		var builder strings.Builder
//		cur := head
//		for cur.Next != nil {
//			builder.WriteString(fmt.Sprintf("%d->", cur.Val))
//			cur = cur.Next
//		}
//		builder.WriteString(fmt.Sprintf("%d", cur.Val))
//		return builder.String()
//	}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy.Next
	total := 0
	for cur != nil {
		total++
		cur = cur.Next
	}

	prev := dummy
	cur = dummy.Next
	for i := 0; i < total-n; i++ {
		prev = cur
		cur = cur.Next
	}
	prev.Next = cur.Next

	return dummy.Next
}

func main() {
	inputs1 := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1},
		[]int{1, 2},
		[]int{1, 2},
	}
	inputs2 := []int{2, 1, 1, 2}
	outputs := [][]int{
		[]int{1, 2, 3, 5},
		[]int{},
		[]int{1},
		[]int{2},
	}

	for pos := 0; pos < len(outputs); pos++ {
		head := createList(inputs1[pos])
		headResult := showList(head)

		a1 := showList(removeNthFromEnd(head, inputs2[pos]))
		a2 := showList(createList(outputs[pos]))

		if a1 != a2 {
			fmt.Println(headResult, a1, a2, a1 == a2)
		}
	}

}
