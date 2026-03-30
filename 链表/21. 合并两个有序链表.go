package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	curr := newHead
	h1 := list1
	h2 := list2
	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			curr.Next = h1
			h1 = h1.Next
		} else {
			curr.Next = h2
			h2 = h2.Next
		}
		curr = curr.Next
	}
	// 连接剩余节点
	if h1 != nil {
		curr.Next = h1
	} else {
		curr.Next = h2
	}
	return newHead.Next
}
