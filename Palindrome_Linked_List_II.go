/*
Leetcode link: https://leetcode.com/problems/reverse-linked-list-ii/
Submission result: Success
Details
Runtime: 3 ms, faster than 35.41% of Go online submissions for Reverse Linked List II.
Memory Usage: 2.1 MB, less than 82.15% of Go online submissions for Reverse Linked List II.
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	if left == 1 {
		l1 := head
		var l2 *ListNode = nil
		if right != LengthListNode(head) {
			l2 = head
			for pos := 1; pos <= right-1; pos++ {
				l2 = l2.Next
			}
			newTail := l2
			l2 = l2.Next
			newTail.Next = nil
		}

		l3 := reverseList(l1)
		head = l3
		for l3.Next != nil {
			l3 = l3.Next
		}
		l3.Next = l2
		return head
	}

	l1 := head
	var l2 *ListNode = nil

	for pos := 1; pos < left-1; pos++ {
		l1 = l1.Next
	}
	newHead := l1.Next

	if right != LengthListNode(head) {
		l2 = newHead
		for pos := left - 1; pos < right-1; pos++ {
			l2 = l2.Next
		}
		newTail := l2
		l2 = l2.Next
		newTail.Next = nil
	}
	l1.Next = reverseList(newHead)
	for l1.Next != nil {
		l1 = l1.Next
	}
	l1.Next = l2

	return head
}

func main() {
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil)))))
	PrintListNode(l)
	PrintListNode(reverseBetween(l, 2, 4))
	l1 := NewListNode(1, NewListNode(2, nil))
	PrintListNode(l1)
	PrintListNode(reverseBetween(l1, 1, 2))
}
