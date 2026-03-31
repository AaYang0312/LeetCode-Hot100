package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}

	pre := dummy
	tail := pre
	for head != nil {
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nex := tail.Next
		// 反转当前组
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex

		pre = tail
		head = tail.Next
	}
	return dummy.Next

}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head
	for prev != tail {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return tail, head
}
