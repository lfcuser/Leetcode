/*
	645. Set Mismatch
	
	You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately,
	due to some error, one of the numbers in s got duplicated to another number in the set,
	which results in repetition of one number and loss of another number.

	You are given an integer array nums representing the data status of this set after the error.
	
	Find the number that occurs twice and the number that is missing and return them in the form of an array.
 */
 
 
package main;

import "fmt"

func findErrorNums(nums []int) []int {
    var res [2]int;
    seen := make(map[int]bool);
    for _, value := range nums {
        if (seen[value] == true) {
            res[0] = value;
        }
        seen[value] = true;
    }
    for i := range len(nums) {
        if (seen[i+1] == false) {
            res[1] = i+1;
            break;
        }
    }
    return res[:];
}

func main() {
	input := []int{1,2,2,4};
    fmt.Println(findErrorNums(input));
    input = []int{1,1};
    fmt.Println(findErrorNums(input));
}

// go run ./golang/SetMismatch/SetMismatch.go

 