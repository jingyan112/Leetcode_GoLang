/*
Leetcode link: https://leetcode.com/problems/merge-in-between-linked-lists/
Submission result: Success
Details
Runtime: 171 ms, faster than 48.48% of Go online submissions for Merge In Between Linked Lists.
Memory Usage: 9.2 MB, less than 6.06% of Go online submissions for Merge In Between Linked Lists.
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

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	breakPoint1 := list1
	for i := 1; i < a; i++ {
		breakPoint1 = breakPoint1.Next
	}
	fmt.Println(breakPoint1.Val)

	breakPoint2 := breakPoint1
	for i := a - 1; i <= b; i++ {
		breakPoint2 = breakPoint2.Next
	}
	fmt.Println(breakPoint2.Val)

	breakPoint1.Next = list2
	newTail := list2
	for newTail.Next != nil {
		newTail = newTail.Next
	}
	newTail.Next = breakPoint2
	return list1
}

func main() {
	list1 := NewListNode(0, NewListNode(1, NewListNode(2, NewListNode(3, NewListNode(4, NewListNode(5, nil))))))
	list2 := NewListNode(100, NewListNode(101, NewListNode(102, nil)))
	PrintListNode(mergeInBetween(list1, 3, 4, list2))
}
