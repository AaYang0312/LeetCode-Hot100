package main

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	tmp := dummy
	// 判断后方是否还有两个节点待交换
	for tmp.Next != nil && tmp.Next.Next != nil {
		node1, node2 := tmp.Next, tmp.Next.Next
		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		tmp = node1
	}
	return dummy.Next
}
