/*
Leetcode link: https://leetcode.com/problems/is-subsequence/
Submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Is Subsequence.
Memory Usage: 2 MB, less than 13.85% of Go online submissions for Is Subsequence.
*/
package main

import "fmt"

func isSubsequence(s string, t string) bool {
	indexS := 0
	indexT := 0
	for indexS < len(s) && indexT < len(t) {
		if s[indexS] == t[indexT] {
			indexS++
			indexT++
		} else {
			indexT++
		}
	}
	if indexS == len(s) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
