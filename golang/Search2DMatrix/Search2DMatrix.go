/*
74 Search a 2D Matrix

You are given an m x n integer matrix matrix with the following two properties:

	Each row is sorted in non-decreasing order.
	The first integer of each row is greater than the last integer of the previous row.

Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Constraints:

	m == matrix.length
	n == matrix[i].length
	1 <= m, n <= 100
	-104 <= matrix[i][j], target <= 104
*/
package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	searchRow := 0
	rowFound := false
	leftRow := 0
	rightRow := m - 1
	for leftRow <= rightRow {
		current := leftRow + (rightRow-leftRow)/2
		if matrix[current][0] == target {
			return true
		}
		if matrix[current][0] > target {
			rightRow = current - 1
		} else {
			if current+1 < m && matrix[current+1][0] <= target {
				leftRow = current + 1
			} else {
				searchRow = current
				rowFound = true
				break
			}
		}
	}

	if rowFound {
		leftCol := 0
		rightCol := n - 1
		for leftCol <= rightCol {
			current := leftCol + (rightCol-leftCol)/2
			if matrix[searchRow][current] == target {
				return true
			}
			if matrix[searchRow][current] > target {
				rightCol = current - 1
			} else {
				leftCol = current + 1
			}
		}
	}
	return false
}

func main() {
	input := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(input, 3))

	input = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix(input, 13))

	input = [][]int{{1}, {3}}
	fmt.Println(searchMatrix(input, 3))
}

// go run ./golang/Search2DMatrix/Search2DMatrix.go
