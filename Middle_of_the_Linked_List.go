/*
Leetcode link: https://leetcode.com/problems/middle-of-the-linked-list/
Submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
Memory Usage: 2 MB, less than 13.91% of Go online submissions for Middle of the Linked List.
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

func printListNode(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, "\t")
		l = l.Next
	}
	fmt.Println()
}

func middlePosListNode(l *ListNode) int {
	nodeNum := 0
	for l != nil {
		l = l.Next
		nodeNum++
	}
	return nodeNum / 2
}

func middleNode(head *ListNode) *ListNode {
	middlePos := middlePosListNode(head)
	l := head
	for i := 0; i < middlePos; i++ {
		l = l.Next
	}
	head = l
	return head
}

func main() {
	l1 := newListNode(7, newListNode(9, newListNode(32, newListNode(56, newListNode(50, nil)))))
	printListNode(middleNode(l1))
	l2 := newListNode(17, newListNode(20, newListNode(35, newListNode(24, newListNode(5, newListNode(6, nil))))))
	printListNode(middleNode(l2))
}
