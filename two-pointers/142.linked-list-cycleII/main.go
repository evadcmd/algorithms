package main

type ListNode struct {
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		slow = slow.Next
		if fast == slow {
			fast = head
			break
		}
	}

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
