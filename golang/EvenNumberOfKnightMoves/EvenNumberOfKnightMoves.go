/*
	3996 Even Number of Knight Moves

	You are given two integer arrays start and target, where each array is of the form [x, y] representing a cell
	on a standard 8 x 8 chessboard.

	Return true if a knight can move from start to target in an even number of moves. Otherwise, return false.

	Note: A valid knight move consists of moving two squares in one direction and one square perpendicular to it.
	The figure below illustrates all eight possible moves from a cell.
*/

package main

import "fmt"

func canReach(start []int, target []int) bool {
	startBlack := (start[0]%2 == 0 && start[1]%2 == 0 || start[0]%2 != 0 && start[1]%2 != 0)
	targetBlack := (target[0]%2 == 0 && target[1]%2 == 0 || target[0]%2 != 0 && target[1]%2 != 0)
	return startBlack == targetBlack
}

func main() {
	input1 := []int{1, 1}
	input2 := []int{2, 2}
	fmt.Println(canReach(input1, input2))
	input1 = []int{4, 5}
	input2 = []int{6, 6}
	fmt.Println(canReach(input1, input2))
}

// go run ./golang/EvenNumberOfKnightMoves/EvenNumberOfKnightMoves.go
