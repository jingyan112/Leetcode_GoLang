/*
Leetcode link: https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
Submission result: Success
Details
Runtime: 291 ms, faster than 13.37% of Go online submissions for Replace Elements with Greatest Element on Right Side.
Memory Usage: 6.5 MB, less than 52.97% of Go online submissions for Replace Elements with Greatest Element on Right Side.
*/

package main

import (
	"fmt"
)

func findGreatest(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}

func replaceElements(arr []int) []int {
	if len(arr) == 1 {
		return []int{-1}
	}
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = findGreatest(arr[i+1 : len(arr)])
	}
	arr[len(arr)-1] = -1
	return arr
}

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
	fmt.Println(replaceElements([]int{400}))
}
