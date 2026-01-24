/*
	9. Palindrome Number

	Given an integer x, return true if x is a , and false otherwise.
 */

package main;

import "fmt"

func isPalindrome(x int) bool {
    copy := x;
    y := 0;
    for (float64(copy) / 10 >= 0.1) {
        y = (y * 10) + (copy % 10);
        copy = copy / 10;
    }
    return y == x;
}

func main() {
    fmt.Println(isPalindrome(121) == true);
    fmt.Println(isPalindrome(-121) == false);
    fmt.Println(isPalindrome(10) == false);
    fmt.Println(isPalindrome(4554) == true);
}

// go run ./golang/PalindromeNumber.go
