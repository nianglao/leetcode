package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	h := head
	for {
		if h == nil {
			break
		}
		size++
		h = h.Next
	}
	t := size - n

	if t <= 0 {
		head = head.Next
	} else {
		h = head
		for {
			if h.Next == nil {
				break
			}
			t--
			if t == 0 {
				if h.Next != nil {
					h.Next = h.Next.Next
				} else {
					h.Next = nil
				}
				break
			}
			h = h.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	PrintList(head)
	PrintList(removeNthFromEnd(head, 5))
}

func PrintList(head *ListNode) {
	for {
		if head == nil {
			break
		}
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
