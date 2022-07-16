package main

import "fmt"

// a -> b -> nil

// prev = nil

// 1. curr = a
//    next = curr.next -> b
//    curr.next = prev -> nil
//    prev = curr -> a
//    curr = next -> b

// 2. curr = b
//    next = curr.next -> c
//    curr.next = prev -> a
//    prev = curr -> b
//    curr = next -> c

// 3. curr = nil

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	currNode := head
	var prevNode *ListNode = nil
	var nextNode *ListNode = nil

	for currNode != nil {
		nextNode = currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}

	return prevNode
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

	return res
}

type data struct {
	input    []int
	expected []int
}

func main() {
	sols := []data{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}

	for _, sol := range sols {
		input := createListNode(sol.input)
		output := reverseList(&input)
		if !solEqual(output, sol.expected) {
			fmt.Println(sol.input, sol.expected)
			showListNode(output)
		}
	}

}
