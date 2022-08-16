/*
Leetcode link: https://leetcode.com/problems/valid-mountain-array/
Submission result: Success
Details
Runtime: 46 ms, faster than 32.03% of Go online submissions for Valid Mountain Array.
Memory Usage: 7 MB, less than 42.86% of Go online submissions for Valid Mountain Array.
*/

package main

import (
	"fmt"
)

func validMountainArray(arr []int) bool {
	indexMaxvalue := 0
	maxValue := 0
	for index, value := range arr {
		if maxValue < value {
			indexMaxvalue = index
			maxValue = value
		}
	}
	if indexMaxvalue == 0 || indexMaxvalue == len(arr)-1 {
		return false
	}
	for i := 0; i < indexMaxvalue; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	for i := indexMaxvalue; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
