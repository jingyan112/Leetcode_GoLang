/*
Leetcode link: https://leetcode.com/problems/valid-palindrome/
Submission result: Success
Details
Runtime: 79 ms, faster than 16.60% of Go online submissions for Valid Palindrome.
Memory Usage: 6.2 MB, less than 19.68% of Go online submissions for Valid Palindrome.
*/
package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	for _, value := range s {
		if !(value >= 'a' && value <= 'z') && !(value >= '0' && value <= '9') {
			if value >= 'A' && value <= 'Z' {
				s = strings.ReplaceAll(s, string(value), string(value|' '))
			} else {
				s = strings.ReplaceAll(s, string(value), "")
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}
