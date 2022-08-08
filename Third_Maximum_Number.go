/*
Leetcode link: https://leetcode.com/problems/third-maximum-number/
submission result: Success
Details
Runtime: 8 ms, faster than 49.55% of Go online submissions for Third Maximum Number.
Memory Usage: 4.5 MB, less than 8.48% of Go online submissions for Third Maximum Number.
*/
package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	numsMap := map[int]bool{}
	var numsNoDuplicate []int
	for _, value := range nums {
		numsMap[value] = false
	}
	for key, _ := range numsMap {
		numsNoDuplicate = append(numsNoDuplicate, key)
	}
	sort.Ints(numsNoDuplicate)
	if len(numsNoDuplicate) < 3 {
		return numsNoDuplicate[len(numsNoDuplicate)-1]
	} else {
		return numsNoDuplicate[len(numsNoDuplicate)-3]
	}
}

func main() {
	fmt.Println(thirdMax([]int{8, 1}))
	fmt.Println(thirdMax([]int{3, 2, 2}))
	fmt.Println(thirdMax([]int{2, 3, 2, 1}))
}
