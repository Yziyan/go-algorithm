// @Author: Ciusyan 9/12/23

package day_7

// https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor1() CQueue {
	return CQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	// 添加时往 inStack 加就行
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	ol := len(this.outStack)
	if ol > 0 {
		// 如果出栈有元素，直接从出栈里面弹出
		val := this.outStack[ol-1]
		this.outStack = this.outStack[:ol-1]

		return val
	}
	il := len(this.inStack)
	if il == 0 {
		return -1
	}

	// 尝试将 inStack 的元素全部弹到 outStack 里
	for i := il - 1; i >= 0; i-- {
		this.outStack = append(this.outStack, this.inStack[i])
	}
	this.inStack = []int{}

	// 再弹出 outStack 的栈顶元素
	val := this.outStack[il-1]
	this.outStack = this.outStack[:il-1]
	return val
}
