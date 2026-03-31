package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	for curr := head; curr != nil; curr = curr.Next.Next {
		curr.Next = &Node{Val: curr.Val, Next: curr.Next}
	}
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next // 指向新创建的节点
		}
	}
	newHead := head.Next
	for curr := head; curr != nil; curr = curr.Next {
		nodeNew := curr.Next
		curr.Next = curr.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return newHead
}
