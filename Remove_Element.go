/*
Leetcode link: https://leetcode.com/problems/remove-element/
Submission result: Success
Details
Runtime: 5 ms, faster than 14.67% of Go online submissions for Remove Element.
Memory Usage: 2.2 MB, less than 14.58% of Go online submissions for Remove Element.
*/
package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	numsNew := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			numsNew = append(numsNew, nums[i])
		}
	}
	for i := 0; i < len(numsNew); i++ {
		nums[i] = numsNew[i]
	}
	return len(numsNew)
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
