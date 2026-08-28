/*
912 Sort an Array

Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Constraints:

    1 <= nums.length <= 5 * 104
    -5 * 104 <= nums[i] <= 5 * 104
*/

package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	length := len(nums)
	if length == 1 {
		return nums
	}

	mid := length / 2
	left := nums[0:mid]
	right := nums[mid:length]
	left = sortArray(left)
	right = sortArray(right)

	i, j := 0, 0
	res := []int{}
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	if i < len(left) {
		res = append(res, left[i])
		i++
	}
	if j < len(right) {
		res = append(res, right[j])
		j++
	}

	return res
}

func main() {
	input := []int{5, 2, 3, 1}
	fmt.Println(sortArray(input))
	input = []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(input))
}

// go run ./golang/SortAnArray/SortAnArray.go
