/*
118 Pascal's Triangle
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it
Constraints: 1 <= numRows <= 30
*/
package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	triangle := [][]int{}

	for i := range numRows {
		newArr := []int{}
		triangle = append(triangle, newArr)
		if i == 0 {
			triangle[i] = append(triangle[i], 1)
			continue
		}
		if i == 1 {
			triangle[i] = append(triangle[i], 1)
			triangle[i] = append(triangle[i], 1)
			continue
		}
		triangle[i] = append(triangle[i], 1)
		for j := 0; j < len(triangle[i-1])-1; j++ {
			triangle[i] = append(triangle[i], triangle[i-1][j]+triangle[i-1][j+1])
		}
		triangle[i] = append(triangle[i], 1)
	}

	return triangle
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}

// go run ./golang/PascalTriangle/PascalTriangle.go
