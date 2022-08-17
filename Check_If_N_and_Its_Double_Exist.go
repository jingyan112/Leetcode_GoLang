/*
Leetcode link: https://leetcode.com/problems/check-if-n-and-its-double-exist/
Submission result: Success
Details
Runtime: 7 ms, faster than 47.97% of Go online submissions for Check If N and Its Double Exist.
Memory Usage: 3.2 MB, less than 68.29% of Go online submissions for Check If N and Its Double Exist.
*/

package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	arrDouble := make([]int, len(arr))
	for index, value := range arr {
		arrDouble[index] = value * 2
	}

	for i := 0; i < len(arrDouble); i++ {
		for j := 0; j < len(arr); j++ {
			if arrDouble[i] == arr[j] && i != j {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 30, 0, 3}))
	fmt.Println(checkIfExist([]int{7, 1, 14, 11}))
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
}
