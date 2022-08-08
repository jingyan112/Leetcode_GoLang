/*
Leetcode link: https://leetcode.com/problems/palindrome-number/
submission result: Success
Details
Runtime: 53 ms, faster than 11.04% of Go online submissions for Palindrome Number.
Memory Usage: 4.8 MB, less than 30.26% of Go online submissions for Palindrome Number.
*/
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		xToString := strconv.Itoa(x)
		for i := 0; i < len(xToString)/2; i++ {
			if xToString[i] != xToString[len(xToString)-1-i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1231))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
