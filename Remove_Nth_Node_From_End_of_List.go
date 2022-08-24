/*
Leetcode link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
Submission result: Success
Details
Runtime: 7 ms, faster than 6.46% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.3 MB, less than 31.66% of Go online submissions for Remove Nth Node From End of List.
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := head
	lengthLinkedList := LengthListNode(l)
	if lengthLinkedList == 1 && n == 1 {
		return nil
	}
	if lengthLinkedList == n {
		return head.Next
	}
	for i := 1; i < lengthLinkedList-n; i++ {
		l = l.Next
	}
	if l.Next.Next == nil {
		l.Next = nil
	} else {
		l.Next = l.Next.Next
	}
	return head
}

func main() {
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil)))))
	PrintListNode(l)
	PrintListNode(removeNthFromEnd(l, 5))
}
