/*
Leetcode link: https://leetcode.com/problems/odd-even-linked-list/
Submission result: Success
Details
Runtime: 4 ms, faster than 70.74% of Go online submissions for Odd Even Linked List.
Memory Usage: 3.3 MB, less than 10.18% of Go online submissions for Odd Even Linked List.
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

func oddEvenList(head *ListNode) *ListNode {
	lengthLinkedList := LengthListNode(head)
	if lengthLinkedList == 0 || lengthLinkedList == 2 {
		return head
	}
	var pos *ListNode = nil
	for pos = head; pos.Next != nil; {
		pos = pos.Next
	}
	Tail := pos

	pos = head
	for index := 1; index < lengthLinkedList; index++ {
		l1 := pos
		if index%2 == 1 {
			newTail := l1.Next
			breakPoint := newTail.Next
			newTail.Next = nil
			Tail.Next = newTail
			Tail = newTail
			l1.Next = breakPoint
			index++
		}
		pos = pos.Next
	}
	return head
}

func main() {
	l1 := NewListNode(1, NewListNode(1, nil))
	PrintListNode(l1)
	PrintListNode(oddEvenList(l1))
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil)))))
	PrintListNode(l)
	PrintListNode(oddEvenList(l))
}
