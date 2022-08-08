/*
Leetcode link: https://leetcode.com/problems/roman-to-integer/
submission result: Success
Details
Runtime: 25 ms, faster than 11.09% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 20.09% of Go online submissions for Roman to Integer.
*/
package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	romanToIntMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	specialRomanToIntMap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	res := 0
	for key, value := range specialRomanToIntMap {
		res = res + strings.Count(s, key)*value
		s = strings.ReplaceAll(s, key, "")
	}
	for i := 0; i < len(s); i++ {
		res = res + romanToIntMap[string(s[i])]
	}
	return res
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
