/*
Leetcode link: https://leetcode.com/problems/linked-list-cycle/
Submission result: Success
Details
Runtime: 11 ms, faster than 63.13% of Go online submissions for Linked List Cycle.
Memory Usage: 4.4 MB, less than 79.85% of Go online submissions for Linked List Cycle.
*/
package main

import "fmt"

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

func hasCycle(head *ListNode) bool {
	l1 := head
	l2 := head
	for l1 != nil && l1.Next != nil && l2 != nil && l2.Next != nil && l2.Next.Next != nil {
		l1 = l1.Next
		l2 = l2.Next.Next
		if l1 == l2 {
			return true
		}
	}
	return false
}

func main() {
	l := newListNode(3, newListNode(2, newListNode(0, newListNode(-4, nil))))
	fmt.Println(hasCycle(l))
	node := l
	pos := 1
	for i := 1; i <= pos; i++ {
		node = node.Next
	}
	fmt.Println(node.Val)

	tail := l
	for tail.Next != nil {
		tail = tail.Next
	}
	fmt.Println(tail.Val)

	tail.Next = node

	fmt.Println(hasCycle(l))
}
