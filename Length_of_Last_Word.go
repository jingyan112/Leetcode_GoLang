/*
Leetcode link: https://leetcode.com/problems/length-of-last-word
submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
Memory Usage: 2.1 MB, less than 87.61% of Go online submissions for Length of Last Word.
*/
package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	sNoWhitespace := strings.Fields(s)
	return len(sNoWhitespace[len(sNoWhitespace)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}
