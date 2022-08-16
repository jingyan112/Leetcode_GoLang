/*
Leetcode link: https://leetcode.com/problems/move-zeroes/
Submission result: Success
Details
Runtime: 37 ms, faster than 48.47% of Go online submissions for Move Zeroes.
Memory Usage: 7.2 MB, less than 37.85% of Go online submissions for Move Zeroes.
*/

package main

import "fmt"

func moveZeroes(nums []int) []int {
	maxLength := len(nums)
	if maxLength == 1 {
		return nums
	}
	for i := 0; i < maxLength; i++ {
		if nums[i] == 0 {
			nums = append(nums[0:i], append(nums[i+1:len(nums)], 0)...)
			i--
			maxLength--
		}
	}
	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
	fmt.Println(moveZeroes([]int{0}))
}
