/*
463 Island Perimeter

You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

Constraints:

    row == grid.length
    col == grid[i].length
    1 <= row, col <= 100
    grid[i][j] is 0 or 1.
    There is exactly one island in grid.
*/

package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	result := 0

	for i := range rows {
		for j := range cols {
			if grid[i][j] == 0 {
				continue
			}
			result += 4
			if j > 0 && grid[i][j-1] == 1 {
				result -= 2
			}
			if i > 0 && grid[i-1][j] == 1 {
				result -= 2
			}
		}
	}

	return result
}

func main() {
	input := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(input), 16)
}

// go run ./golang/IslandPerimeter/IslandPerimeter.go
