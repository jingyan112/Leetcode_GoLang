/*
Leetcode link: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
Submission result: Success
Details
Runtime: 63 ms, faster than 74.56% of Go online submissions for Find All Numbers Disappeared in an Array.
Memory Usage: 8.3 MB, less than 41.81% of Go online submissions for Find All Numbers Disappeared in an Array.
*/

package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	integerArray := make([]int, len(nums))
	for i := 1; i <= len(nums); i++ {
		integerArray[i-1] = i
	}

	numsMap := make(map[int]bool)
	for _, value := range nums {
		numsMap[value] = false
	}

	missedArray := make([]int, 0)
	for _, value := range integerArray {
		if _, ok := numsMap[value]; !ok {
			missedArray = append(missedArray, value)
		}
	}

	return missedArray
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
