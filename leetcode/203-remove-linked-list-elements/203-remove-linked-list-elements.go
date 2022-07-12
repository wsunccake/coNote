package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeElements(head *ListNode, val int) *ListNode {
// 	currNode := head
// 	var prevNode *ListNode = nil
// 	var nextNode *ListNode = nil
// 	var resNode *ListNode = head
// 	for currNode != nil {
// 		nextNode = currNode.Next
// 		if currNode.Val == val {
// 			if prevNode != nil {
// 				prevNode.Next = nextNode
// 			} else {
// 				resNode = nextNode
// 			}
// 			currNode.Next = nil
// 		} else {
// 			prevNode = currNode
// 		}
// 		currNode = nextNode
// 	}
// 	return resNode
// }

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	currNode := head
	prevNode := &ListNode{0, nil}
	prevNode.Next = currNode

	for currNode != nil {
		nextNode := currNode.Next

		if currNode.Val == val {
			prevNode.Next = nextNode
		} else {
			prevNode = currNode
		}
		currNode = nextNode
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	return head
}

func createListNode(s []int) ListNode {
	if len(s) == 0 {
		return ListNode{}
	}

	headNode := ListNode{s[0], nil}
	currNode := &headNode
	for i := 1; i < len(s); i++ {
		(*currNode).Next = &ListNode{s[i], nil}
		currNode = (*currNode).Next
	}
	return headNode
}

func showListNode(node *ListNode) {
	for node != nil {
		next := node.Next
		node = next
	}
}

func solEqual(node *ListNode, sols []int) bool {
	res := true
	currNode := node
	for _, sol := range sols {
		if currNode.Val != sol {
			res = false
			break
		}
		currNode = currNode.Next
	}

	if len(sols) == 0 && node != nil {
		res = false
	}

	return res
}

type data struct {
	input    []int
	value    int
	expected []int
}

func main() {
	sols := []data{
		{[]int{1, 2, 6, 3, 4, 5, 6}, 6, []int{1, 2, 3, 4, 5}},
		{[]int{}, 1, []int{}},
		{[]int{7, 7, 7, 7}, 7, []int{}},
		{[]int{6}, 6, []int{}},
	}

	for _, sol := range sols {
		input := createListNode(sol.input)
		output := removeElements(&input, sol.value)
		// showListNode(output)
		if !solEqual(output, sol.expected) {
			fmt.Println(sol.input, sol.expected)
			// showListNode(output)
		}
	}

}
