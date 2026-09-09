/*
3516 Find Closest Person

You are given three integers x, y, and z, representing the positions of three people on a number line:

	x is the position of Person 1.
	y is the position of Person 2.
	z is the position of Person 3, who does not move.

Both Person 1 and Person 2 move toward Person 3 at the same speed.

Determine which person reaches Person 3 first:

	Return 1 if Person 1 arrives first.
	Return 2 if Person 2 arrives first.
	Return 0 if both arrive at the same time.

Return the result accordingly.

Constraints:

	1 <= x, y, z <= 100
*/
package main

import (
	"fmt"
	"math"
)

func findClosest(x int, y int, z int) int {
	xToZ := math.Abs(float64(x - z))
	yToZ := math.Abs(float64(y - z))
	if xToZ < yToZ {
		return 1
	} else if yToZ < xToZ {
		return 2
	}
	return 0
}

func main() {
	fmt.Println(findClosest(2, 7, 4), 1)
	fmt.Println(findClosest(2, 5, 6), 2)
	fmt.Println(findClosest(1, 5, 3), 0)
}

// go run ./golang/FindClosestPerson/FindClosestPerson.go
