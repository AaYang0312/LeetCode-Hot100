package main

func hasCycle(head *ListNode) bool {
	// 龟兔赛跑
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	if slow == fast {
		return true
	}
	return false
}
