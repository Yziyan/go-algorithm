// @Author: Ciusyan 3/17/24

package phase_1

type MinStack struct {
	head *node
}

type node struct {
	val    int
	minVal int
	Next   *node
}

func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &node{val: val, minVal: val}
		return
	}

	// 说明以前已经有节点了
	minVal := val
	if this.head.minVal < minVal {
		// 说明栈顶的还要小
		minVal = this.head.minVal
	}
	// 新建节点接入头部
	n := &node{val: val, minVal: minVal}
	n.Next = this.head
	this.head = n
}

func (this *MinStack) Pop() {
	if this.head != nil {
		this.head = this.head.Next
	}
}

func (this *MinStack) Top() int {
	if this.head == nil {
		return -1
	}
	return this.head.val
}

func (this *MinStack) GetMin() int {
	if this.head == nil {
		return -1
	}
	return this.head.minVal
}
