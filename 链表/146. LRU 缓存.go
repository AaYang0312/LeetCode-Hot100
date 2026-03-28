package main

type Node struct {
	k, v      int
	next, pre *Node
}
type LRUCache struct {
	cap        int
	head, tail *Node
	m          map[int]*Node
}

func Constructor(capacity int) LRUCache {
	h, t := &Node{}, &Node{}
	h.next = t
	t.pre = h
	return LRUCache{capacity, h, t, make(map[int]*Node)}
}
func (this *LRUCache) moveToHead(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	this.addToHead(node)
}
func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	temp := this.m
	if node, ok := temp[key]; ok {
		this.moveToHead(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	temp := this.m
	if node, ok := temp[key]; ok {
		node.v = value
		this.moveToHead(node)
		return
	}
	if len(temp) >= this.cap {
		deleteNode := this.tail.pre
		delete(temp, deleteNode.k)
		deleteNode.pre.next = this.tail
		this.tail.pre = deleteNode.pre

	}
	newNode := &Node{key, value, nil, nil}
	this.addToHead(newNode)
	temp[key] = newNode
}
