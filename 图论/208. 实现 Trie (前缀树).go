package main

type Trie struct {
	root *Node
}
type Node struct {
	son [26]*Node
	end bool
}

func Constructor() Trie {
	return Trie{&Node{}}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil {
			node := &Node{}
			cur.son[c] = node
		}
		cur = cur.son[c]
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		c -= 'a'
		if cur.son[c] == nil {
			return false
		}
		cur = cur.son[c]
	}
	if cur.end == true {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		c -= 'a'
		if cur.son[c] == nil {
			return false
		}
		cur = cur.son[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
