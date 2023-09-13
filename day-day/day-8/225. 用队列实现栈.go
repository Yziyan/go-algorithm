// @Author: Ciusyan 9/13/23

package day_8

// https://leetcode.cn/problems/implement-stack-using-queues/

type queue []int

func (s *queue) Poll() int {
	if len(*s) == 0 {
		return 0
	}
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}

func (s *queue) Offer(x int) {
	*s = append(*s, x)
}

type MyStack struct {
	q1   *queue
	q2   *queue
	size int
}

func Constructor() MyStack {
	return MyStack{
		q1: &queue{},
		q2: &queue{},
	}
}

func (this *MyStack) Push(x int) {
	// 加入的时候，不管怎么样，都从 q1 加
	this.q1.Offer(x)
	this.size++
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}

	// 先去操作
	res := this.top()
	this.q1 = &queue{}
	// 更换后需要换指向
	this.q1, this.q2 = this.q2, this.q1
	this.size--

	return res
}

func (this *MyStack) top() int {
	if this.Empty() {
		return 0
	}
	l := len(*this.q1)
	// 出队 len - 1 个，只剩最后一个
	for l > 1 {
		// 将 q1 的元素弹到 q2 中去
		this.q2.Offer(this.q1.Poll())
		l--
	}

	return (*this.q1)[0]
}

func (this *MyStack) Top() int {
	return this.top()
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}
