/*
390 Elimination Game

You have a list arr of all integers in the range [1, n] sorted in a strictly increasing order. Apply the following algorithm on arr:

	Starting from left to right, remove the first number and every other number afterward until you reach the end of the list.
	Repeat the previous step again, but this time from right to left, remove the rightmost number and every other number from the remaining numbers.
	Keep repeating the steps again, alternating left to right and right to left, until a single number remains.

Given the integer n, return the last number that remains in arr.

Constraints:

	1 <= n <= 109
*/
package main

import (
	"fmt"
)

func lastRemaining(n int) int {
	if n == 1 {
		return n
	}
	left := 2
	right := n
	if n%2 > 0 {
		right = n - 1
	}
	leftToRight := false
	counter := 2
	length := right / 2
	for left < right {
		isOddLen := length%2 > 0
		if leftToRight {
			if isOddLen {
				right -= counter
			}
			left += counter
		} else {
			if isOddLen {
				left += counter
			}
			right -= counter
		}
		counter *= 2
		length /= 2
		leftToRight = !leftToRight
	}
	return left
}

func main() {
	fmt.Println(lastRemaining(9), 6)
	fmt.Println(lastRemaining(1), 1)
	fmt.Println(lastRemaining(10), 8)
	fmt.Println(lastRemaining(12), 6)
	fmt.Println(lastRemaining(16), 6)
	fmt.Println(lastRemaining(24), 14)
}

// go run ./golang/EliminationGame/EliminationGame.go
