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
