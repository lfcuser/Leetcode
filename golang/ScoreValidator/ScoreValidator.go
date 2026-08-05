/*
3921 Score Validator

You are given a string array events.

Initially, score = 0 and counter = 0. Each element in events is one of the following:

    "0", "1", "2", "3", "4", "6": Add that value to the total score.
    "W": Increase the counter by 1. No score is added.
    "WD": Add 1 to the total score.
    "NB": Add 1 to the total score.

Process the array from left to right. Stop processing when either:

    All elements in events have been processed, or
    The counter becomes 10.

Return an integer array [score, counter], where:

    score is the final total score.
    counter is the final counter value.

Constraints:

1 <= events.length <= 1000
events[i] is one of "0", "1", "2", "3", "4", "6", "W", "WD", or "NB".
*/

package main

import (
	"fmt"
	"strconv"
)

func scoreValidator(events []string) []int {
	score := 0
	counter := 0
	for _, value := range events {
		if value == "W" {
			counter++
		} else if value == "NB" || value == "WD" {
			score++
		} else {
			num, _ := strconv.Atoi(value)
			score += num
		}
		if counter >= 10 {
			break
		}
	}
	return []int{score, counter}
}

func main() {
	input := []string{"1", "4", "W", "6", "WD"}
	fmt.Println(scoreValidator(input))
}

// go run ./golang/ScoreValidator/ScoreValidator.go
