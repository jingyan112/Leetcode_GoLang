/*
Leetcode link: https://leetcode.com/problems/squares-of-a-sorted-array/
Submission result: Success
Details
Runtime: 59 ms, faster than 38.05% of Go online submissions for Squares of a Sorted Array.
Memory Usage: 7.4 MB, less than 38.84% of Go online submissions for Squares of a Sorted Array.
*/

package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
