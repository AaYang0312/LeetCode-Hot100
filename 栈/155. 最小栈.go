package main

import "math"

type pair struct{ val, preMin int }

type MinStack []pair

func Constructor() MinStack {
	// 这里的 0 写成任意数都可以，反正用不到
	return MinStack{{0, math.MaxInt}} // 栈底哨兵
}

func (st *MinStack) Push(val int) {
	*st = append(*st, pair{val, min(st.GetMin(), val)})
}

func (st *MinStack) Pop() {
	*st = (*st)[:len(*st)-1]
}

func (st MinStack) Top() int {
	return st[len(st)-1].val
}

func (st MinStack) GetMin() int {
	return st[len(st)-1].preMin
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
