/*
Leetcode link: https://leetcode.com/problems/majority-element/
Submission result: Success
Details
Runtime: 29 ms, faster than 44.35% of Go online submissions for Majority Element.
Memory Usage: 6.2 MB, less than 68.97% of Go online submissions for Majority Element.
*/
package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	numsMap := map[int]int{}
	for _, value := range nums {
		numsMap[value] = 0
	}
	for _, value := range nums {
		numsMap[value] = numsMap[value] + 1
	}
	var res int
	for key, value := range numsMap {
		if value > len(nums)/2 {
			res = key
		}
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
