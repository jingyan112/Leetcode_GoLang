/*
Leetcode link: https://leetcode.com/problems/reverse-linked-list/
Submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
Memory Usage: 2.6 MB, less than 76.26% of Go online submissions for Reverse Linked List.
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

func insertFromHead(node1 *ListNode, node2 *ListNode) *ListNode {
	node2.Next = node1
	return node2
}

func reverseList(head *ListNode) *ListNode {
	var newListHead *ListNode = nil
	var l1 *ListNode = nil
	for ptr := head; ptr != nil; ptr = l1 {
		l1 = ptr.Next
		ptr.Next = nil
		if newListHead == nil {
			newListHead = ptr
		} else {
			newListHead = insertFromHead(newListHead, ptr)
		}
	}
	return newListHead
}

func main() {
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, NewListNode(6, nil))))))
	PrintListNode(l)
	PrintListNode(reverseList(l))
}
