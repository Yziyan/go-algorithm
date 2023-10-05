// @Author: Ciusyan 10/5/23

package day_19

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack 简单使用切片当做一个栈
type Stack []*TreeNode

func NewStack() Stack {
	return make([]*TreeNode, 0)
}

func (s *Stack) Push(val *TreeNode) {
	// 添加到末尾即可
	*s = append(*s, val)
}

func (s *Stack) Pop() *TreeNode {
	// 从末尾弹出，并且将末尾元素删除即可
	last := len(*s) - 1
	res := (*s)[last]
	*s = (*s)[:last]

	return res
}

func (s *Stack) Size() int {
	return len(*s)
}
