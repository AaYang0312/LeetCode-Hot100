package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 找中间节点，断开前一个
	head2 := middleNode(head)
	// 分治
	head1 := sortList(head)
	head2 = sortList(head2)
	// 合并
	return mergeLists(head1, head2)
}
func middleNode(head *ListNode) *ListNode {
	slow, pre, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil // 断开
	return slow
}
func mergeLists(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			curr.Next = head1
			head1 = head1.Next
		} else {
			curr.Next = head2
			head2 = head2.Next
		}
		curr = curr.Next
	}
	if head1 != nil {
		curr.Next = head1
	}
	if head2 != nil {
		curr.Next = head2
	}
	return dummy.Next
}
