/*
Leetcode link: https://leetcode.com/problems/remove-linked-list-elements/
Submission result: Success
Details
Runtime: 9 ms, faster than 69.94% of Go online submissions for Remove Linked List Elements.
Memory Usage: 4.7 MB, less than 19.64% of Go online submissions for Remove Linked List Elements.
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

func removeElements(head *ListNode, val int) *ListNode {
	var l *ListNode
	l = head
	for l != nil && l.Next != nil {
		if l.Val == val {
			head = head.Next
		}
		if l.Next.Val == val {
			l.Next = l.Next.Next
		} else {
			l = l.Next
		}
	}
	if head != nil && head.Next == nil && head.Val == val {
		head = nil
	}
	return head
}

func main() {
	l1 := newListNode(1, newListNode(2, newListNode(6, newListNode(3, newListNode(4, newListNode(5, newListNode(6, nil)))))))
	//printListNode(l1)
	printListNode(removeElements(l1, 6))
	l2 := newListNode(7, newListNode(7, newListNode(7, nil)))
	//printListNode(l2)
	printListNode(removeElements(l2, 7))
}
