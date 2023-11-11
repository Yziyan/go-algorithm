// @Author: Ciusyan 11/11/23

package day_24

// ReverseStackUsingRecursive 使用递归的方式逆序栈，要求不能使用如何额外的数据机构。
func ReverseStackUsingRecursive(stack *Stack) {
	if stack.Size() == 0 {
		return
	}

	// 先移除栈底元素
	button := removeAndButton(stack)
	// 然后对剩下的栈逆序
	ReverseStackUsingRecursive(stack)
	// 逆序后，再将栈底元素放置到栈顶
	stack.Push(button)
}

// 将栈底元素删除，并返回
func removeAndButton(stack *Stack) int {
	// 先弹出来一个
	pop := stack.Pop()
	if stack.Size() == 0 {
		// 如果弹出来后就没有元素了，直接返回就好
		return pop
	}

	// 否则递归此过程，获得最后栈底元素，并且将上面的盖下来
	last := removeAndButton(stack)
	// 但是在返回前，需要将原先弹出来的一个给加回去
	stack.Push(pop)

	return last
}

// Stack 模拟一个栈
type Stack []int

func NewStack() *Stack {
	res := make(Stack, 0)
	return &res
}

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[:last]
	return res
}

func (s *Stack) Size() int {
	return len(*s)
}
