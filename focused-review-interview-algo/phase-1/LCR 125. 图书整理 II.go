// @Author: Ciusyan 3/17/24

package phase_1

type CQueue struct {
	inStack  []int
	outStack []int
}

func CQueueConstructor() CQueue {
	return CQueue{
		inStack:  make([]int, 0, 10),
		outStack: make([]int, 0, 10),
	}
}

func (this *CQueue) AppendTail(value int) {
	// 加入无脑加入入栈
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	// 先看看能否从出栈弹出元素
	oL := len(this.outStack)
	if oL > 0 {
		// 说明能从出栈直接弹出元素
		res := this.outStack[oL-1]
		this.outStack = this.outStack[:oL-1]
		return res
	}

	// 否则看看入栈有没有元素
	iL := len(this.inStack)
	if iL == 0 {
		// 说明没有元素
		return -1
	}

	// 说明有，现将其全部加入 outStack
	for iL > 1 {
		iL--
		this.outStack = append(this.outStack, this.inStack[iL])
	}
	// 还有最后一个元素，就是要返回的
	res := this.inStack[0]
	// 但是要将入栈还原
	this.inStack = make([]int, 0, iL)

	return res
}
