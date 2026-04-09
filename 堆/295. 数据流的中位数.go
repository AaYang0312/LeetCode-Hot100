package main

import "container/heap"

type IntHeap []int
type minHeap struct {
	IntHeap
}
type maxHeap struct {
	IntHeap
}

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// minHeap
func (mh minHeap) Less(i, j int) bool {
	return mh.IntHeap[i] < mh.IntHeap[j]
}

// maxHeap
func (mh maxHeap) Less(i, j int) bool {
	return mh.IntHeap[i] > mh.IntHeap[j]
}

type MedianFinder struct {
	left  *maxHeap
	right *minHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &maxHeap{},
		right: &minHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == this.right.Len() {
		heap.Push(this.right, num)
		heap.Push(this.left, heap.Pop(this.right))
	} else {
		heap.Push(this.left, num)
		heap.Push(this.right, heap.Pop(this.left))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64(this.right.IntHeap[0]+this.left.IntHeap[0]) / 2
	} else {
		return float64(this.left.IntHeap[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
