/*
Leetcode link: https://leetcode.com/problems/reorder-list/
Submission result: Success
Details
Runtime: 365 ms, faster than 92.31% of Go online submissions for Merge Nodes in Between Zeros.
Memory Usage: 12.8 MB, less than 93.27% of Go online submissions for Merge Nodes in Between Zeros.
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

func mergeNodes(head *ListNode) *ListNode {
	var newHead *ListNode = NewListNode(0, nil)
	newSumListNode := newHead
	pos1 := head
	sumBetween := 0
	for pos2 := pos1.Next; pos2.Next != nil; pos2 = pos2.Next {
		if pos2.Next.Val == 0 {
			if pos2.Next.Next != nil {
				newSumListNode.Next = NewListNode(0, nil)
			}
			sumBetween = sumBetween + pos2.Val
			pos1 = pos2.Next
			newSumListNode.Val = sumBetween
			sumBetween = 0
			newSumListNode = newSumListNode.Next
		} else {
			sumBetween = sumBetween + pos2.Val
		}
	}
	return newHead
}

func main() {
	l := NewListNode(0, NewListNode(3, NewListNode(1, NewListNode(0, NewListNode(4, NewListNode(5, NewListNode(2, NewListNode(0, nil))))))))
	PrintListNode(l)
	PrintListNode(mergeNodes(l))
	l1 := NewListNode(0, NewListNode(1, NewListNode(0, NewListNode(3, NewListNode(0, NewListNode(2, NewListNode(2, NewListNode(0, nil))))))))
	PrintListNode(l1)
	PrintListNode(mergeNodes(l1))
}
