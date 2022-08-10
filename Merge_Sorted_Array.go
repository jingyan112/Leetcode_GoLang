/*
Leetcode link: https://leetcode.com/problems/merge-sorted-array/
Submission result: Success
Details
Runtime: 6 ms, faster than 12.42% of Go online submissions for Merge Sorted Array.
Memory Usage: 2.5 MB, less than 17.83% of Go online submissions for Merge Sorted Array.
*/
package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i <= len(nums1)-1; i++ {
		nums1[i] = nums2[i-m]
	}
	if len(nums2) != 0 && len(nums1) != 0 && m != 0 && nums1[m-1] > nums2[0] {
		sort.Ints(nums1)
	}
	fmt.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}
