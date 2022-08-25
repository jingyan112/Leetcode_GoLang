/*
Leetcode link: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
Submission result: Success
Details
Runtime: 375 ms, faster than 77.57% of Go online submissions for Delete the Middle Node of a Linked List.
Memory Usage: 13.1 MB, less than 14.95% of Go online submissions for Delete the Middle Node of a Linked List.
*/
package main

import "fmt"

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

func deleteMiddle(head *ListNode) *ListNode {
	lengthLinkedList := LengthListNode(head)
	if lengthLinkedList == 1 {
		return nil
	}
	l := head
	for i := 1; i < lengthLinkedList/2; i++ {
		l = l.Next
	}
	if l.Next.Next != nil {
		l.Next = l.Next.Next
	} else {
		l.Next = nil
	}
	return head
}

func main() {
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil)))))
	PrintListNode(l)
	PrintListNode(deleteMiddle(l))
	l1 := NewListNode(1, NewListNode(2, nil))
	PrintListNode(l1)
	PrintListNode(deleteMiddle(l1))
}
