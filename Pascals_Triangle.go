/*
Leetcode link: https://leetcode.com/problems/pascals-triangle/
Submission result: Success
Details
Runtime: 3 ms, faster than 20.05% of Go online submissions for Pascal's Triangle.
Memory Usage: 2 MB, less than 87.23% of Go online submissions for Pascal's Triangle.
*/
package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
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
	return res
}

func main() {
	fmt.Println(generate(6))
	//fmt.Println(generate(1))
}
