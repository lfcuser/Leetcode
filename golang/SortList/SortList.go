/*
148 Sort List

Given the head of a linked list, return the list after sorting it in ascending order.

Input: head = [4,2,1,3]
Output: [1,2,3,4]

Constraints:

	The number of nodes in the list is in the range [0, 5 * 104].
	-105 <= Node.val <= 105
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	arr := listToArray(head)

	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		newHead := &ListNode{Val: arr[0]}
		return newHead
	}

	arr = mergeSort(arr)
	newHead := &ListNode{}
	return fillList(arr, 0, newHead)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}

	mid := length / 2
	left := arr[0:mid]
	right := arr[mid:length]

	left = mergeSort(left)
	right = mergeSort(right)

	result := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func fillList(arr []int, i int, head *ListNode) *ListNode {
	if i >= len(arr) {
		return nil
	}

	head.Val = arr[i]
	next := &ListNode{}
	head.Next = fillList(arr, i+1, next)
	return head
}

func listToArray(head *ListNode) []int {
	arr := []int{}

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}

func main() {
	input := []int{4, 2, 1, 3}
	head := &ListNode{}
	list := fillList(input, 0, head)
	fmt.Println(listToArray(list))
	fmt.Println(listToArray(sortList(list)))
}

// go run ./golang/SortList/SortList.go
