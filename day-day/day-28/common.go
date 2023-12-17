// @Author: Ciusyan 12/17/23

package day_28

// Stack 模拟一个栈
type Stack []int

func NewStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[0:last]
	return res
}

func (s *Stack) Peek() int {
	last := len(*s) - 1
	return (*s)[last]
}

func (s *Stack) Size() int {
	return len(*s)
}

// StackList 模拟一个栈
type StackList [][]int

func NewStackList() StackList {
	return make(StackList, 0)
}

func (s *StackList) Push(v []int) {
	*s = append(*s, v)
}

func (s *StackList) Pop() []int {
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[0:last]
	return res
}

func (s *StackList) Peek() []int {
	last := len(*s) - 1
	return (*s)[last]
}

func (s *StackList) Size() int {
	return len(*s)
}
