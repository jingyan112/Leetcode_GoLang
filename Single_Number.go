/*
Leetcode link: https://leetcode.com/problems/single-number/
Submission result: Success
Details
Runtime: 54 ms, faster than 6.32% of Go online submissions for Single Number.
Memory Usage: 7.3 MB, less than 8.64% of Go online submissions for Single Number.
*/
package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	numsMap := make(map[int]int)
	for _, value := range nums {
		numsMap[value] = 0
	}
	for _, value := range nums {
		numsMap[value]++
	}
	fmt.Println(numsMap)
	for key, value := range numsMap {
		if value == 1 {
			return key
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}
