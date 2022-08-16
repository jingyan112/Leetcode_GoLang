/*
Leetcode link: https://leetcode.com/problems/rotate-array/
Submission result:
*/
package main

import "fmt"

func rotate(nums []int, k int) []int {
	for i := 1; i <= k; i++ {
		nums = append(nums[len(nums)-1:len(nums)], nums[0:len(nums)-1]...)
	}
	return nums
}

func main() {
	fmt.Println(rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	fmt.Println(rotate([]int{-1, -100, 3, 99}, 2))
}
