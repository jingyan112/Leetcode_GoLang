/*
Leetcode link: https://leetcode.com/problems/rotate-list/
Submission result: Success
Details
Runtime: 2 ms, faster than 70.26% of Go online submissions for Rotate List.
Memory Usage: 2.6 MB, less than 82.03% of Go online submissions for Rotate List.
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(Val int, Next *ListNode) *ListNode {
	var n ListNode
	n.Val = Val
	n.Next = Next
	return &n
}

func PrintListNode(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, "\t")
		l = l.Next
	}
	fmt.Println()
}

func LengthListNode(l *ListNode) int {
	length := 0
	for l != nil {
		l = l.Next
		length++
	}
	return length
}

func rotateRight(head *ListNode, k int) *ListNode {
	if LengthListNode(head) == 0 {
		return head
	} else {
		k = k % LengthListNode(head)
		if k == 0 {
			return head
		}
	}
	fmt.Println("Rotate: ", k, ", Length: ", LengthListNode(head))
	index := head
	SecondHalfHead := head
	for i := 0; i < LengthListNode(head)-k-1; i++ {
		index = index.Next
	}
	SecondHalfTail := index
	FirstHalfHead := index.Next
	FirstHalfTail := index
	for FirstHalfTail.Next != nil {
		FirstHalfTail = FirstHalfTail.Next
	}
	SecondHalfTail.Next = nil
	FirstHalfTail.Next = SecondHalfHead
	return FirstHalfHead
}

func main() {
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, NewListNode(6, nil))))))
	PrintListNode(l)
	PrintListNode(rotateRight(l, 8))
}
