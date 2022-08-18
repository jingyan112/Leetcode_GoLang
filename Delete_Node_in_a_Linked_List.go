/*
Leetcode link: https://leetcode.com/problems/delete-node-in-a-linked-list/
Submission result: Success
Details
Runtime: 6 ms, faster than 34.83% of Go online submissions for Delete Node in a Linked List.
Memory Usage: 2.9 MB, less than 71.39% of Go online submissions for Delete Node in a Linked List.
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

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	l := newListNode(4, newListNode(5, newListNode(1, newListNode(9, nil))))
	printListNode(l)
	node := l
	pos := 2
	for i := 1; i <= pos; i++ {
		node = node.Next
	}
	deleteNode(node)
	printListNode(l)
}
