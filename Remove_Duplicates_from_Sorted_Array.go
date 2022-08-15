/*
Leetcode link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
Submission result: Success
Details
Runtime: 13 ms, faster than 55.97% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 6.4 MB, less than 5.59% of Go online submissions for Remove Duplicates from Sorted Array.
*/
package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(nums []int) int {
	numsMap := make(map[int]bool)
	for _, value := range nums {
		numsMap[value] = false
	}

	tmpArray := make([]int, 0)
	for key, _ := range numsMap {
		tmpArray = append(tmpArray, key)
	}
	sort.Ints(tmpArray)

	for index, value := range tmpArray {
		nums[index] = value
	}

	return len(numsMap)
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
