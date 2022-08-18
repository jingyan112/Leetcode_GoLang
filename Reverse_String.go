/*
Leetcode link: https://leetcode.com/problems/reverse-string/
Submission result: Success
Details
Runtime: 37 ms, faster than 79.88% of Go online submissions for Reverse String.
Memory Usage: 6.6 MB, less than 72.79% of Go online submissions for Reverse String.
*/

package main

import "fmt"

func reverseString(s []byte) {
	fmt.Println(string(s))
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = tmp
	}
	fmt.Println(string(s))
}

func main() {
	reverseString([]byte("hello"))
	reverseString([]byte("Hannah"))
}
