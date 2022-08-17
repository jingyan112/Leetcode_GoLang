/*
Leetcode link: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
Submission result: Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Duplicates from Sorted List.
Memory Usage: 3.2 MB, less than 40.03% of Go online submissions for Remove Duplicates from Sorted List.
*/

package main

import (
	"fmt"
)

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

func deleteDuplicates(head *ListNode) *ListNode {
	l := head
	for l != nil && l.Next != nil {
		if l.Val != l.Next.Val {
			l = l.Next
		} else {
			if l.Next.Next == nil {
				l.Next = nil
			} else {
				l.Next = l.Next.Next
			}
		}
	}
	return head
}

func main() {
	l1 := newListNode(1, newListNode(1, newListNode(2, nil)))
	printListNode(l1)
	printListNode(deleteDuplicates(l1))
	l2 := newListNode(1, newListNode(1, newListNode(2, newListNode(2, newListNode(3, newListNode(3, nil))))))
	printListNode(l2)
	printListNode(deleteDuplicates(l2))
}
