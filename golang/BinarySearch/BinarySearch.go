// 704 Binary Search

package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		current := left + (right-left)/2
		if nums[current] == target {
			return current
		}
		if nums[current] < target {
			left = current + 1
		} else {
			right = current - 1
		}
	}
	return -1
}

func main() {
	input := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(input, 9))
}

// go run ./golang/BinarySearch/BinarySearch.go
