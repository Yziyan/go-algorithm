// @Author: Ciusyan 10/13/23

package day_20

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func NewQueue() Queue {
	return make([]*TreeNode, 0)
}

func (q *Queue) Offer(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Poll() *TreeNode {
	node := (*q)[0]
	*q = (*q)[1:]

	return node
}

func (q *Queue) Size() int {
	return len(*q)
}

type Stack []int

func NewStack() Stack {
	return make([]int, 0)
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

func (s *Stack) Top() int {
	last := len(*s) - 1
	res := (*s)[last]
	return res
}

func (s *Stack) Size() int {
	return len(*s)
}
