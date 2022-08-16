/*
Leetcode link: https://leetcode.com/problems/sort-array-by-parity/
Submission result: Success
Details
Runtime: 125 ms, faster than 5.13% of Go online submissions for Sort Array By Parity.
Memory Usage: 9.5 MB, less than 5.56% of Go online submissions for Sort Array By Parity.
*/

package main

import (
	"fmt"
)

func sortArrayByParity(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			tmp := nums[i]
			tmp1 := nums[0:i]
			tmp2 := nums[i+1 : len(nums)]
			nums = append([]int{tmp}, append(tmp1, tmp2...)...)
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
}
