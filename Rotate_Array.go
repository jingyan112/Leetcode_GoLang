/*
Leetcode link: https://leetcode.com/problems/rotate-array/
Submission result: Success
Details
Runtime: 59 ms, faster than 38.21% of Go online submissions for Rotate Array.
Memory Usage: 8.9 MB, less than 28.46% of Go online submissions for Rotate Array.
*/
package main

import "fmt"

func rotate(nums []int, k int) []int {
	var tmp []int = make([]int, len(nums))
	copy(tmp, nums)
	k = k % len(tmp)
	tmp = append(tmp[len(tmp)-k:len(tmp)], tmp[0:len(tmp)-k]...)
	fmt.Println(tmp)
	copy(nums, tmp)
	return nums
}

func main() {
	fmt.Println(rotate([]int{1, 2}, 5))
}
