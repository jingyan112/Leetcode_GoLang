/*
Leetcode link: https://leetcode.com/problems/height-checker/
Submission result: Success
Details
Runtime: 4 ms, faster than 35.98% of Go online submissions for Height Checker.
Memory Usage: 2.4 MB, less than 9.76% of Go online submissions for Height Checker.
*/

package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	var heightsCopy []int = make([]int, len(heights))
	copy(heightsCopy, heights)
	sort.Ints(heights)
	count := 0
	fmt.Println(heightsCopy, heights)
	for i := 0; i < len(heights); i++ {
		if heightsCopy[i] != heights[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))
}
