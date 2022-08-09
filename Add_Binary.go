/*
Leetcode link: https://leetcode.com/problems/add-binary/
Submission result: Success
Details
Runtime: 3 ms, faster than 51.12% of Go online submissions for Add Binary.
Memory Usage: 3.2 MB, less than 8.26% of Go online submissions for Add Binary.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	c := map[int]int{}
	for i := len(a); i > 0; i-- {
		if string(a[i-1])+string(b[i-1]) == "11" {
			c[i] = 2
		}
		if string(a[i-1])+string(b[i-1]) == "10" || string(a[i-1])+string(b[i-1]) == "01" {
			c[i] = 1
		}
		if string(a[i-1])+string(b[i-1]) == "00" {
			c[i] = 0
		}
	}
	c[0] = 0
	for i := len(a); i > 0; i-- {
		if c[i] >= 2 {
			c[i] = c[i] - 2
			c[i-1]++
		}
	}
	res := ""
	flag := 0
	if c[0] == 0 {
		flag = 1
	}
	for i := flag; i <= len(a); i++ {
		res = res + strconv.Itoa(c[i])
	}
	if res != "0" {
		for {
			if strings.HasPrefix(res, "1") {
				break
			} else {
				res = res[1:len(res)]
			}
		}
	}
	return res
}

func main() {
	fmt.Println(addBinary("100", "110010"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("1111", "1111"))
}
