/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, tail, cur *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur = l1
			l1 = l1.Next
		} else {
			cur = l2
			l2 = l2.Next
		}
		if head == nil {
			head = cur
		}
		if tail == nil {
			tail = cur
		} else {
			tail.Next = cur
			tail = cur
		}
	}

	if l1 != nil {
		tail.Next = l1
	}

	if l2 != nil {
		tail.Next = l2
	}
	return head
}
