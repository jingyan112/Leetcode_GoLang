/*
Leetcode link: https://leetcode.com/problems/reorder-list/
Submission result: Success
Details
Runtime: 83 ms, faster than 5.07% of Go online submissions for Reorder List.
Memory Usage: 6.1 MB, less than 31.04% of Go online submissions for Reorder List.
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

func ReverseList(head *ListNode) *ListNode {
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

func reorderList(head *ListNode) *ListNode {
	lengthLinkedList := LengthListNode(head)
	if lengthLinkedList == 0 || lengthLinkedList == 1 || lengthLinkedList == 2 {
		return head
	}
	pos := head
	for i := 0; i < lengthLinkedList/2; i++ {
		pos = pos.Next
	}
	fmt.Println("Tail is: ", pos.Val)
	newHead := pos.Next
	pos.Next = nil
	newLinkedList := ReverseList(newHead)
	PrintListNode(newLinkedList)
	fmt.Println("Length of new linked list is: ", LengthListNode(newLinkedList))

	lengthNewLinkedList := LengthListNode(newLinkedList)
	if lengthNewLinkedList == 1 {
		pos1 := head
		pos2 := head.Next
		pos1.Next = newLinkedList
		newLinkedList.Next = pos2
		return head
	}

	pos1 := head
	pos2 := newLinkedList
	for i := 1; i <= lengthNewLinkedList; i++ {
		nextPos1 := pos1.Next
		nextPos2 := pos2.Next
		pos1.Next = pos2
		pos2.Next = nextPos1
		pos1 = pos2.Next
		pos2 = nextPos2
	}
	return head
}

func main() {
	//l1 := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, nil))))
	//PrintListNode(l1)
	//PrintListNode(reorderList(l1))
	l := NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil)))))
	PrintListNode(l)
	PrintListNode(reorderList(l))
}
