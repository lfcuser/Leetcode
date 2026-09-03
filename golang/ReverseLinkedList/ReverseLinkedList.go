/*
206 Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Constraints:

	The number of nodes in the list is the range [0, 5000].
	-5000 <= Node.val <= 5000
*/
package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return reverse(nil, head)
}

func reverse(prev *ListNode, current *ListNode) *ListNode {
	next := current.Next
	current.Next = prev
	if next == nil {
		return current
	}
	return reverse(current, next)
}

func main() {
	first := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	first.Next = second

	before, _ := json.MarshalIndent(first, "", "    ")
	fmt.Println(string(before))

	after, _ := json.MarshalIndent(reverseList(first), "", "    ")
	fmt.Println(string(after))
}

// go run ./golang/ReverseLinkedList/ReverseLinkedList.go
