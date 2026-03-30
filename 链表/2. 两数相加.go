package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	curr := newHead
	h1 := l1
	h2 := l2
	carry := 0 // 进位
	for h1 != nil || h2 != nil || carry > 0 {
		sum := carry
		if h1 != nil {
			sum += h1.Val
			h1 = h1.Next
		}
		if h2 != nil {
			sum += h2.Val
			h2 = h2.Next
		}
		curr.Val = sum % 10
		carry = sum / 10
		// 后续还有节点或进位，创建新节点
		if h1 != nil || h2 != nil || carry > 0 {
			newNode := &ListNode{}
			curr.Next = newNode
			curr = curr.Next
		}
	}
	return newHead
}
