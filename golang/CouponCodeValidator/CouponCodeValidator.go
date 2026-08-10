/*
3606 Coupon Code Validator

You are given three arrays of length n that describe the properties of n coupons: code, businessLine, and isActive. The ith coupon has:

    code[i]: a string representing the coupon identifier.
    businessLine[i]: a string denoting the business category of the coupon.
    isActive[i]: a boolean indicating whether the coupon is currently active.

A coupon is considered valid if all of the following conditions hold:

    code[i] is non-empty and consists only of alphanumeric characters (a-z, A-Z, 0-9) and underscores (_).
    businessLine[i] is one of the following four categories: "electronics", "grocery", "pharmacy", "restaurant".
    isActive[i] is true.

Return an array of the codes of all valid coupons, sorted first by their businessLine in the order: "electronics", "grocery", "pharmacy", "restaurant", and then by code in lexicographical (ascending) order within each category.



Example 1:

Input: code = ["SAVE20","","PHARMA5","SAVE@20"], businessLine = ["restaurant","grocery","pharmacy","restaurant"], isActive = [true,true,true,true]

Output: ["PHARMA5","SAVE20"]

Explanation:

    First coupon is valid.
    Second coupon has empty code (invalid).
    Third coupon is valid.
    Fourth coupon has special character @ (invalid).

Example 2:

Input: code = ["GROCERY15","ELECTRONICS_50","DISCOUNT10"], businessLine = ["grocery","electronics","invalid"], isActive = [false,true,true]

Output: ["ELECTRONICS_50"]

Explanation:

    First coupon is inactive (invalid).
    Second coupon is valid.
    Third coupon has invalid business line (invalid).



Constraints:

    n == code.length == businessLine.length == isActive.length
    1 <= n <= 100
    0 <= code[i].length, businessLine[i].length <= 100
    code[i] and businessLine[i] consist of printable ASCII characters.
    isActive[i] is either true or false.
*/

package main

import (
	"fmt"
	"regexp"
	"slices"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	electronics := []string{}
	grocery := []string{}
	pharmacy := []string{}
	restaurant := []string{}

	re := regexp.MustCompile(`^[A-Za-z0-9_]+$`)

	length := len(code)
	for i := range length {
		if !re.MatchString(code[i]) || !isActive[i] {
			continue
		}
		switch businessLine[i] {
		case "electronics":
			electronics = append(electronics, code[i])
		case "grocery":
			grocery = append(grocery, code[i])
		case "pharmacy":
			pharmacy = append(pharmacy, code[i])
		case "restaurant":
			restaurant = append(restaurant, code[i])
		}
	}

	slices.Sort(electronics[:])
	slices.Sort(grocery[:])
	slices.Sort(pharmacy[:])
	slices.Sort(restaurant[:])
	result := slices.Concat(electronics, grocery, pharmacy, restaurant)

	return result
}

func main() {
	input1 := []string{"SAVE20", "", "PHARMA5", "SAVE@20"}
	input2 := []string{"restaurant", "grocery", "pharmacy", "restaurant"}
	input3 := []bool{true, true, true, true}
	fmt.Println(validateCoupons(input1, input2, input3))
}

// go run ./golang/CouponCodeValidator/CouponCodeValidator.go
