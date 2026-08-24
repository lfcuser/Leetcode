/*
2409 Count Days Spent Together

Alice and Bob are traveling to Rome for separate business meetings.

You are given 4 strings arriveAlice, leaveAlice, arriveBob, and leaveBob. Alice will be in the city from the dates arriveAlice to leaveAlice (inclusive), while Bob will be in the city from the dates arriveBob to leaveBob (inclusive). Each will be a 5-character string in the format "MM-DD", corresponding to the month and day of the date.

Return the total number of days that Alice and Bob are in Rome together.

You can assume that all dates occur in the same calendar year, which is not a leap year. Note that the number of days per month can be represented as: [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31].

Constraints:

	All dates are provided in the format "MM-DD".
	Alice and Bob's arrival dates are earlier than or equal to their leaving dates.
	The given dates are valid dates of a non-leap year.
*/
package main

import (
	"fmt"
	"strconv"
)

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	monthAliceArrive, dayAliceArrive := parse(arriveAlice)
	monthAliceLeave, dayAliceLeave := parse(leaveAlice)
	monthBobArrive, dayBobArrive := parse(arriveBob)
	monthBobLeave, dayBobLeave := parse(leaveBob)

	monthStart, dayStart := lateArrive(monthAliceArrive, monthBobArrive, dayAliceArrive, dayBobArrive)
	monthEnd, dayEnd := earlyLeave(monthAliceLeave, monthBobLeave, dayAliceLeave, dayBobLeave)

	if monthEnd >= monthStart {
		return countDaysBetween(monthStart, monthEnd, dayStart, dayEnd)
	}
	return 0
}

func parse(input string) (int, int) {
	month, _ := strconv.Atoi(string(input[0]) + string(input[1]))
	day, _ := strconv.Atoi(string(input[3]) + string(input[4]))
	return month, day
}

func lateArrive(month1 int, month2 int, day1 int, day2 int) (int, int) {
	if month1 > month2 {
		return month1, day1
	} else if month1 < month2 {
		return month2, day2
	}
	return month1, max(day1, day2)
}

func earlyLeave(month1 int, month2 int, day1 int, day2 int) (int, int) {
	if month1 < month2 {
		return month1, day1
	} else if month1 > month2 {
		return month2, day2
	}
	return month1, min(day1, day2)
}

func countDaysBetween(monthStart int, monthEnd int, dayStart int, dayEnd int) int {
	months := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	daysMonth := 0
	start := monthStart - 1
	for start < monthEnd {
		daysMonth += months[start]
		start++
	}
	res := daysMonth - dayStart - (months[monthEnd-1] - dayEnd) + 1
	return max(res, 0)
}

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"), 3)
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31"), 0)
	fmt.Println(countDaysTogether("09-01", "10-19", "06-19", "10-20"), 49)
	fmt.Println(countDaysTogether("08-06", "12-08", "02-04", "09-01"), 27)
	fmt.Println(countDaysTogether("03-05", "07-14", "04-14", "09-21"), 92)
	fmt.Println(countDaysTogether("12-26", "12-27", "08-21", "12-17"), 0)
}

// go run ./golang/CountDaysSpentTogether/CountDaysSpentTogether.go
