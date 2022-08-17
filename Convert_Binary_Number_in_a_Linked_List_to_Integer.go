/*
Leetcode link: https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer
Submission result: Success
Details
Runtime: 4 ms, faster than 27.27% of Go online submissions for Convert Binary Number in a Linked List to Integer.
Memory Usage: 2.1 MB, less than 29.55% of Go online submissions for Convert Binary Number in a Linked List to Integer.
*/

package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(Val int, Next *ListNode) *ListNode {
	var n ListNode
	n.Val = Val
	n.Next = Next
	return &n
}

func printListNode(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, "\t")
		l = l.Next
	}
	fmt.Println()
}

func lengthListNode(l *ListNode) int {
	length := 0
	for l != nil {
		l = l.Next
		length++
	}
	return length
}

func getDecimalValue(head *ListNode) int {
	decimalValue := 0
	l := head
	for i := lengthListNode(head) - 1; i >= 0; i-- {
		decimalValue = decimalValue + l.Val*int(math.Pow(2, float64(i)))
		l = l.Next
	}
	return decimalValue
}

func main() {
	l := newListNode(1, newListNode(0, newListNode(1, nil)))
	fmt.Println(getDecimalValue(l))
	fmt.Println(getDecimalValue(newListNode(0, nil)))
}
