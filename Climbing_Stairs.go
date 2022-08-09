/*
Leetcode link: https://leetcode.com/problems/climbing-stairs/
Submission result: Success
Details
Runtime: 9 ms, faster than 19.60% of Go online submissions for Climbing Stairs.
Memory Usage: 4.7 MB, less than 8.57% of Go online submissions for Climbing Stairs.
*/
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func calc(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, calc(n.Sub(x, n)))
}

func BigIntToStr(bigInt *big.Int) string {
	return fmt.Sprintf("%v", bigInt)
}

func climbStairs(n int) int {
	var count *big.Int = big.NewInt(0)
	var waysArray [][]int
	for i := 0; i <= n; i++ {
		for j := 0; j <= n/2; j++ {
			if n == i+j*2 {
				waysArray = append(waysArray, []int{i, j})
			}
		}
	}
	for _, value := range waysArray {
		var tmp *big.Int = big.NewInt(1)
		numerator := tmp.Add(new(big.Int).SetInt64(int64(value[0])), new(big.Int).SetInt64(int64(value[1])))
		numerator = calc(numerator)
		denominator := tmp.Mul(calc(new(big.Int).SetInt64(int64(value[0]))), calc(new(big.Int).SetInt64(int64(value[1]))))
		numerator = tmp.Div(numerator, denominator)
		count = tmp.Add(count, numerator)
	}
	i, _ := strconv.Atoi(BigIntToStr(count))
	return i
}

func main() {
	fmt.Println(climbStairs(35))
}
