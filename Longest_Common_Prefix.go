/*
Leetcode link: https://leetcode.com/problems/longest-common-prefix/
submission result: Success
Details
Runtime: 5 ms, faster than 27.45% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.9 MB, less than 9.51% of Go online submissions for Longest Common Prefix.
*/
package main

import (
	"fmt"
	"strings"
)

func isprefix(strs []string, prefix string) bool {
	for _, value := range strs {
		if !strings.HasPrefix(value, prefix) {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	maxLength := len(strs[0])
	for _, value := range strs {
		if maxLength > len(value) {
			maxLength = len(value)
		}
	}
	var possiblePrefixArray []string
	for i := maxLength; i > 0; i-- {
		possiblePrefix := ""
		for j := 0; j < i; j++ {
			possiblePrefix = possiblePrefix + string(strs[0][j])
		}
		possiblePrefixArray = append(possiblePrefixArray, possiblePrefix)
	}
	for _, value := range possiblePrefixArray {
		if isprefix(strs, value) {
			return value
		}
	}
	return ""
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "dog", "do"}))
}
