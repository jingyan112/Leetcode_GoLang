/*
Leetcode link: https://leetcode.com/problems/add-two-numbers/
submission result: Success
Details
Runtime: 13 ms, faster than 66.60% of Go online submissions for Add Two Numbers.
Memory Usage: 4.6 MB, less than 45.26% of Go online submissions for Add Two Numbers.
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(value int, node *ListNode) *ListNode {
	var n ListNode
	n.Val = value
	n.Next = node
	return &n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := NewListNode(l1.Val+l2.Val, nil)
	l3 := head
	for l1.Next != nil || l2.Next != nil {
		if l1.Next != nil {
			l1 = l1.Next
			if l2.Next != nil {
				l2 = l2.Next
				l3.Next = NewListNode(l1.Val+l2.Val, nil)
			} else {
				l3.Next = NewListNode(l1.Val, nil)
			}
		} else {
			if l2.Next != nil {
				l2 = l2.Next
				l3.Next = NewListNode(l2.Val, nil)
			}
		}
		//fmt.Println(l3.Val)
		l3 = l3.Next
	}
	l4 := head
	for l4 != nil {
		if l4.Val >= 10 {
			l4.Val = l4.Val - 10
			if l4.Next == nil {
				l4.Next = NewListNode(1, nil)
			} else {
				l4.Next.Val++
			}
		}
		l4 = l4.Next
	}
	return head
}

func main() {
	l1 := NewListNode(9, NewListNode(9, NewListNode(9, NewListNode(9, NewListNode(9, NewListNode(9, NewListNode(9, nil)))))))
	l2 := NewListNode(9, NewListNode(9, NewListNode(9, NewListNode(9, nil))))
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
