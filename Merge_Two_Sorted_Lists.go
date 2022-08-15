/*
Leetcode link: https://leetcode.com/problems/merge-two-sorted-lists/
Submission result: Success
Details
Runtime: 8 ms, faster than 13.18% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.7 MB, less than 9.13% of Go online submissions for Merge Two Sorted Lists.
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

func printListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var l3 *ListNode = nil
	var head *ListNode
	for list1 != nil || list2 != nil {
		if l3 == nil {
			l3 = newListNode(0, nil)
			head = l3
		} else {
			l3.Next = newListNode(0, nil)
			l3 = l3.Next
		}
		if list1 == nil {
			l3.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			l3.Val = list1.Val
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				l3.Val = list1.Val
				list1 = list1.Next
			} else {
				l3.Val = list2.Val
				list2 = list2.Next
			}
		}
	}
	return head
}

func main() {
	l1 := newListNode(1, newListNode(2, newListNode(4, nil)))
	l2 := newListNode(1, newListNode(3, newListNode(4, nil)))
	printListNode(mergeTwoLists(l1, l2))
}
