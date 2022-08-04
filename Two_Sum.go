/*
Leetcode link: https://leetcode.com/problems/two-sum/
Submission result: Success
Details
Runtime: 59 ms, faster than 14.27% of Go online submissions for Two Sum.
Memory Usage: 3.6 MB, less than 94.98% of Go online submissions for Two Sum.
*/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var IndexArray = make([]int, 2)
	var FirstIndex, SecondIndex int
	for FirstIndex = 0; FirstIndex < len(nums)-1; FirstIndex++ {
		for SecondIndex = len(nums) - 1; SecondIndex > FirstIndex; SecondIndex-- {
			if nums[FirstIndex]+nums[SecondIndex] == target {
				IndexArray = []int{FirstIndex, SecondIndex}
			}
		}
	}
	return IndexArray
}

func main() {
	var nums = make([]int, 10)
	nums = []int{2, 1, 1, 1, 1, 1, 1, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}
