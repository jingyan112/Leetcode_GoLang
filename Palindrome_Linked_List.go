/*
Leetcode link: https://leetcode.com/problems/palindrome-linked-list/
Submission result: Success
Details
Runtime: 213 ms, faster than 68.68% of Go online submissions for Palindrome Linked List.
Memory Usage: 10.3 MB, less than 70.27% of Go online submissions for Palindrome Linked List.
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

func isPalindrome(head *ListNode) bool {
	firstLinkedList := head
	l := head
	lengthLinkedList := LengthListNode(head)
	if lengthLinkedList == 1 || lengthLinkedList == 0 {
		return true
	}
	for i := 0; i < lengthLinkedList/2-1; i++ {
		l = l.Next
	}
	if lengthLinkedList%2 == 1 {
		l = l.Next
	}
	secondHalfHead := l.Next
	l.Next = nil
	secondLinkedList := reverseList(secondHalfHead)
	for i := 0; i <= lengthLinkedList/2-1; i++ {
		if firstLinkedList.Val != secondLinkedList.Val {
			return false
		}
		firstLinkedList = firstLinkedList.Next
		secondLinkedList = secondLinkedList.Next
	}
	return true
}

func main() {
	//l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(2, NewListNode(1, nil)))))
	//fmt.Println(isPalindrome(l))
	l1 := NewListNode(1, nil)
	fmt.Println(isPalindrome(l1))
}
