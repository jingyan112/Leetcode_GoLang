/*
Leetcode link: https://leetcode.com/problems/search-insert-position/
Submission result: Success
Details
Runtime: 4 ms, faster than 70.65% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 76.97% of Go online submissions for Search Insert Position.
*/
package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	var pos int
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for min <= max {
		fmt.Println("Before: ", min, max, pos)
		pos = (min + max) / 2
		if pos == min {
			if target == nums[pos] {
				return pos
			}
			return pos + 1
		}
		if nums[pos] < target {
			min = pos
		} else {
			max = pos
		}
		fmt.Println("After: ", min, max, pos)
	}
	return pos
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5}, 4))
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3}, 2))
}
