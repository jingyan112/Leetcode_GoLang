/*
Leetcode link: https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
Submission result: Success
Details
Runtime: 6 ms, faster than 57.81% of Go online submissions for Find Numbers with Even Number of Digits.
Memory Usage: 3.1 MB, less than 21.26% of Go online submissions for Find Numbers with Even Number of Digits.
*/

package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	count := 0
	for _, value := range nums {
		if len(strconv.Itoa(value))%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}
