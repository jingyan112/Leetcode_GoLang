/*
Leetcode link: https://leetcode.com/problems/valid-parentheses/
submission result: Success
Details
Runtime: 3 ms, faster than 35.15% of Go online submissions for Valid Parentheses.
Memory Usage: 7.6 MB, less than 5.25% of Go online submissions for Valid Parentheses.
*/
package main

import (
	"fmt"
	"strings"
)

func hasPairBrackets(s string) string {
	for _, bracket := range []string{"()", "[]", "{}"} {
		if strings.Contains(s, bracket) {
			return hasPairBrackets(strings.ReplaceAll(s, bracket, ""))
		}
	}
	return s
}

func isValid(s string) bool {
	if hasPairBrackets(s) == "" {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("[({(())}[()])]"))
}
