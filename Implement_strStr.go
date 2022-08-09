/*
Leetcode link: https://leetcode.com/problems/implement-strstr/
submission result: Success
Details
Runtime: 1 ms, faster than 51.93% of Go online submissions for Implement strStr().
Memory Usage: 1.9 MB, less than 93.85% of Go online submissions for Implement strStr().
*/
package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if strings.Contains(haystack, needle) {
		for i := 0; i < len(haystack); i++ {
			if string(haystack[i:i+len(needle)]) == needle {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("mississippi", "issip"))
}
