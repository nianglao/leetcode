package leetcode

// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			if head == nil {
				head = l1
				tail = l1
			} else {
				tail.Next = l1
				tail = l1
			}
			l1 = l1.Next
		} else {
			if head == nil {
				head = l2
				tail = l2
			} else {
				tail.Next = l2
				tail = l2
			}
			l2 = l2.Next
		}
	}
	for {
		if l1 == nil {
			break
		}
		tail.Next = l1
		tail = l1
		l1 = l1.Next
	}

	for {
		if l2 == nil {
			break
		}
		tail.Next = l2
		tail = l2
		l2 = l2.Next
	}
	return head
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

func main() {

	var l1, l2 *ListNode
	l1 = &ListNode{Val: 1, Next: &ListNode{2, nil}}
	l2 = &ListNode{Val: 1, Next: &ListNode{3, nil}}
	PrintList(mergeTwoLists(l1, l2))

}
