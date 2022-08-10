/*
Leetcode link: https://leetcode.com/problems/plus-one/
Submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2.1 MB, less than 64.16% of Go online submissions for Plus One.
*/
package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	index := len(digits) - 1
	for index >= 1 {
		if digits[index] == 10 {
			digits[index] = 0
			digits[index-1]++
		}
		index--
	}
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9}))
	fmt.Println(plusOne([]int{1, 2, 3}))
}
