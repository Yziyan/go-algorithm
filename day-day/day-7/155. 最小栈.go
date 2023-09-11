// @Author: Ciusyan 9/12/23

package day_7

// https://leetcode.cn/problems/min-stack/

type MinStack struct {
	Head *node
}

type node struct {
	Val    int
	MinVal int
	Next   *node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.Head == nil {
		this.Head = &node{Val: val, MinVal: val}
		return
	}

	// 来到这里不是 nil
	min := this.Head.MinVal
	if val < min {
		min = val
	}

	// 新建节点，并插入头部
	this.Head = &node{Val: val, MinVal: min, Next: this.Head}
}

func (this *MinStack) Pop() {
	this.Head = this.Head.Next
}

func (this *MinStack) Top() int {
	if n := this.Head; n != nil {
		return n.Val
	}

	return 0
}

func (this *MinStack) GetMin() int {
	if n := this.Head; n != nil {
		return n.MinVal
	}

	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
