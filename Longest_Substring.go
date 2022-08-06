/*
Leetcode link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
submission result: Success
Details
Runtime: 244 ms, faster than 18.78% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 9.2 MB, less than 5.60% of Go online submissions for Longest Substring Without Repeating Characters.
*/
package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	for Start := 0; Start < len(s); Start++ {
		subString := string(s[Start])
		for Index := Start + 1; Index < len(s); Index++ {
			if !strings.Contains(subString, string(s[Index])) {
				subString = subString + string(s[Index])
			} else {
				break
			}
		}
		if len(subString) > maxLength {
			maxLength = len(subString)
		}
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcdabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}
