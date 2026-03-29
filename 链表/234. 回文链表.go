package main

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针，找到慢指针在中点或中点右节点
	var prev *ListNode = nil
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	var head2 *ListNode = nil
	// 奇数中点不需要回退，可能导致其他bug，最后比较时一方走到nil就退出了，多余中间节点不影响
	// 反转链表
	for slow != nil {
		next := slow.Next
		slow.Next = head2
		head2 = slow
		slow = next
	}
	// 验证回文
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		} else {
			head = head.Next
			head2 = head2.Next
		}
	}
	return true
}
