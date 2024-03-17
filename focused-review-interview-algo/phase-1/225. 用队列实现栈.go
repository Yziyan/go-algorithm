// @Author: Ciusyan 3/17/24

package phase_1

type MyStack struct {
	inQueue  []int
	outQueue []int
	size     int
}

func Constructor() MyStack {
	return MyStack{
		inQueue:  make([]int, 0, 10),
		outQueue: make([]int, 0, 10),
		size:     0,
	}
}

func (this *MyStack) Push(x int) {
	// 加入元素直接无脑从入队加即可
	this.inQueue = append(this.inQueue, x)
	this.size++
}

func (this *MyStack) Pop() int {
	if this.size == 0 {
		return -1
	}
	this.size--
	// 这就好了
	return this.change()
}

// 置换 inQueue 和 outQueue，并返回 inQueue 中的最后一个元素
func (this *MyStack) change() int {
	iL := len(this.inQueue)
	for i := 0; i < iL-1; i++ {
		// 将 inQueue 的元素，除了队尾的，其余全部加入 outQueue 中
		this.outQueue = append(this.outQueue, this.inQueue[i])
	}
	// 这是 inQueue 中最后一个元素
	res := this.inQueue[iL-1]
	// 在返回前，需要先交换两个队列的应用
	this.inQueue, this.outQueue = this.outQueue, this.inQueue
	return res
}

func (this *MyStack) Top() int {
	if this.size == 0 {
		return -1
	}
	res := this.change()
	// 但是上面相当于把栈顶元素弹出了，还要加回去
	this.inQueue = append(this.inQueue, res)

	return res
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}
