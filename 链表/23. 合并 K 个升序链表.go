package main

// 分治 迭代
func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:m/2])
	right := mergeKLists(lists[m/2:])
	return mergeTwoLists(left, right)
}
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	newHead := &ListNode{}
	head1, head2 := list1, list2
	curr := newHead
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
	} else if head2 != nil {
		curr.Next = head2
	}
	return newHead.Next
}
