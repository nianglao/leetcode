package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for {
		if head == nil {
			break
		}
		fmt.Printf("%v\n", head.Val)
		head = head.Next
	}
}
