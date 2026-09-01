/*
141 Linked List Cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Constraints:
    The number of the nodes in the list is in the range [0, 104].
    -105 <= Node.val <= 105
    pos is -1 or a valid index in the linked-list.
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	objectsMap := make(map[*ListNode]bool)

	for head != nil {
		_, exist := objectsMap[head]
		if exist {
			return true
		}
		objectsMap[head] = true
		head = head.Next
	}

	return false
}

func main() {
	first := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	first.Next = second
	second.Next = first
	fmt.Println(hasCycle(first))

	first = &ListNode{Val: 1}
	second = &ListNode{Val: 2}
	first.Next = second
	fmt.Println(hasCycle(first))
}

// go run ./golang/LinkedListCycle/LinkedListCycle.go
