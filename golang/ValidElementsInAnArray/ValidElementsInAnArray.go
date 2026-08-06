/*
3912 Valid Elements in an Array
You are given an integer array nums.

An element nums[i] is considered valid if it satisfies AT LEAST one of the following conditions:

    It is strictly greater than every element to its left.
    It is strictly greater than every element to its right.

The first and last elements are always valid.

Return an array of all valid elements in the same order as they appear in nums.

Example 1:

Input: nums = [1,2,4,2,3,2]

Output: [1,2,4,3,2]

Explanation:

    nums[0] and nums[5] are always valid.
    nums[1] and nums[2] are strictly greater than every element to their left.
    nums[4] is strictly greater than every element to its right.
    Thus, the answer is [1, 2, 4, 3, 2].

Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

package main

import (
	"fmt"
	"slices"
)

func findValidElements(nums []int) []int {
	length := len(nums)
	result := []int{}
	left := 0

	for i := range length {
		if nums[i] > left || i == length-1 {
			result = append(result, nums[i])
			left = nums[i]
			continue
		}
		rights := nums[i+1 : length]
		if nums[i] > slices.Max(rights) {
			result = append(result, nums[i])
		}
	}

	return result
}

func main() {
	input := []int{1, 2, 4, 2, 3, 2}
	fmt.Println(findValidElements(input))
}

// go run ./golang/ValidElementsInAnArray/ValidElementsInAnArray.go
