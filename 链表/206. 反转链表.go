package main

// 迭代
// func reverseList(head *ListNode) *ListNode {
// 	var pre *ListNode
// 	curr := head
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = pre
// 		pre = curr
// 		curr = next

// 	}
// 	return pre
// }

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
