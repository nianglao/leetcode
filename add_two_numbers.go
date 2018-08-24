package leetcode

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Printf("%v->", l.Val)
		l = l.Next
	}
	fmt.Printf("\n")
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := l1
	inc := 0
	for l1.Next != nil && l2.Next != nil {
		tmp := l1.Val
		l1.Val = (l1.Val + l2.Val + inc) % 10
		inc = (tmp + l2.Val + inc) / 10
		l1 = l1.Next
		l2 = l2.Next

	}
	if l1.Next == nil && l2.Next == nil {
		tmp := l1.Val
		l1.Val = (l1.Val + l2.Val + inc) % 10
		inc = (tmp + l2.Val + inc) / 10
		if inc > 0 {
			l1.Next = &ListNode{inc, nil}
		}
	} else {
		if l1.Next == nil {
			l1.Next = l2.Next
		}
		tmp := l1.Val
		l1.Val = (l1.Val + l2.Val + inc) % 10
		inc = (tmp + l2.Val + inc) / 10
		l1 = l1.Next
		for l1.Next != nil {
			tmp := l1.Val
			l1.Val = (l1.Val + inc) % 10
			inc = (tmp + inc) / 10
			l1 = l1.Next
		}
		tmp = l1.Val
		l1.Val = (l1.Val + inc) % 10
		inc = (tmp + inc) / 10
		if inc > 0 {
			l1.Next = &ListNode{inc, nil}
		}
	}

	return ret

}

func main() {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}
	printList(l1)
	printList(l2)
	printList(addTwoNumbers(l1, l2))
}
