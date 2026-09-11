/*
3232 Find if Digit Game Can Be Won

You are given an array of positive integers nums.
Alice and Bob are playing a game. In the game, Alice can choose either all single-digit numbers or
all double-digit numbers from nums, and the rest of the numbers are given to Bob.
Alice wins if the sum of her numbers is strictly greater than the sum of Bob's numbers.
Return true if Alice can win this game, otherwise, return false.

Constraints:

	1 <= nums.length <= 100
	1 <= nums[i] <= 99
*/
package main

import "fmt"

func canAliceWin(nums []int) bool {
	singleDigits := 0
	doubleDigits := 0
	for _, val := range nums {
		if val/10 > 0 {
			doubleDigits += val
		} else {
			singleDigits += val
		}
	}
	if singleDigits == doubleDigits {
		return false
	}
	return true
}

func main() {
	input := []int{1, 2, 3, 4, 10}
	fmt.Println(canAliceWin(input), false)
	input = []int{1, 2, 3, 4, 5, 14}
	fmt.Println(canAliceWin(input), true)
}

// go run ./golang/FindIfDigitGameCanBeWon/FindIfDigitGameCanBeWon.go
