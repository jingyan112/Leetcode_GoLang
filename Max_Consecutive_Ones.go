/*
Leetcode link: https://leetcode.com/problems/max-consecutive-ones/
Submission result: Success
Details
Runtime: 145 ms, faster than 5.59% of Go online submissions for Max Consecutive Ones.
Memory Usage: 7.4 MB, less than 13.97% of Go online submissions for Max Consecutive Ones.
*/

package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			continue
		} else {
			index++
		}
	}
	if nums[len(nums)-1] == 1 {
		index++
	}

	numsMap := make(map[int]int)
	for i := 1; i <= index; i++ {
		numsMap[i] = 0
	}

	index = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			numsMap[index]++
		} else {
			index++
			continue
		}
	}

	fmt.Println(numsMap)

	maxConsecutive := 0
	for _, value := range numsMap {
		if value > maxConsecutive {
			maxConsecutive = value
		}
	}

	return maxConsecutive
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1, 0, 1}))
}
