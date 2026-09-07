/*
69 Sqrt(x)

Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.
For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
Constraints:

	0 <= x <= 231 - 1
*/
package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	res := 2
	for true {
		buf := res * res
		if buf == x {
			break
		}
		if buf > x {
			res--
			break
		}
		res++
	}
	return res
}

func main() {
	fmt.Println(mySqrt(4), 2)
	fmt.Println(mySqrt(8), 2)
	fmt.Println(mySqrt(9), 3)
}

// go run ./golang/Sqrt/Sqrt.go
