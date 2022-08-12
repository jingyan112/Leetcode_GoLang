/*
Leetcode link: https://leetcode.com/problems/pascals-triangle-ii/
Submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle II.
Memory Usage: 2.1 MB, less than 33.15% of Go online submissions for Pascal's Triangle II.
*/
package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
	}
	res[0] = []int{1}
	if numRows >= 2 {
		res[1] = []int{1, 1}
		for i := 2; i < len(res); i++ {
			res[i][0] = 1
			res[i][i] = 1
			for j := 1; j <= i-1; j++ {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res[rowIndex]
}

func main() {
	fmt.Println(getRow(1))
}
