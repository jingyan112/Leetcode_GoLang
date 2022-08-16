/*
Leetcode link: https://leetcode.com/problems/duplicate-zeros/
Submission result: Success
Details
Runtime: 8 ms, faster than 61.54% of Go online submissions for Duplicate Zeros.
Memory Usage: 3.4 MB, less than 23.89% of Go online submissions for Duplicate Zeros.
*/

package main

import (
	"fmt"
)

func duplicateZeros(arr []int) []int {
	count := 0
	for _, value := range arr {
		if value == 0 {
			count++
		}
	}
	if count == 0 {
		return arr
	}
	arrNew := make([]int, len(arr)+count)
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arrNew[index] = arr[i]
			index = index + 1
		} else {
			arrNew[index] = 0
			arrNew[index+1] = 0
			index = index + 2

		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = arrNew[i]
	}
	return arr
}

func main() {
	fmt.Println(duplicateZeros([]int{1, 0, 2, 3, 0, 4, 5, 0}))
	fmt.Println(duplicateZeros([]int{1, 2, 3}))
}
