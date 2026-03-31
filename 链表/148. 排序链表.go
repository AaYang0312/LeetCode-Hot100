package main

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}

	// 子串长度
	for subLength := 1; subLength < length; subLength <<= 1 {
		// 每次sublength变化都要初始化
		prev, curr := dummyHead, dummyHead.Next
		for curr != nil {
			head1 := curr
			// 找到第一个子串尾部
			for i := 1; i < subLength && curr.Next != nil; i++ {
				curr = curr.Next
			}
			head2 := curr.Next
			// 切断串间连接
			curr.Next = nil
			curr = head2
			// 找到第2个子串尾部
			for i := 1; i < subLength && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}
			var nextHead *ListNode
			if curr != nil {
				nextHead = curr.Next
				curr.Next = nil
			}
			prev.Next = Merge(head1, head2)
			// prev 放到队尾
			for prev.Next != nil {
				prev = prev.Next
			}
			curr = nextHead // nexthead存在则为下一子串开头，否则为nil

		}

	}
	return dummyHead.Next

}
func Merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp1, temp2 := head1, head2
	curr := dummy
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			curr.Next = temp1
			temp1 = temp1.Next
			curr = curr.Next
		} else {
			curr.Next = temp2
			temp2 = temp2.Next
			curr = curr.Next
		}
	}
	if temp1 != nil {
		curr.Next = temp1
	} else if temp2 != nil {
		curr.Next = temp2
	}
	return dummy.Next
}
