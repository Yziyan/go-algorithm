// @Author: Ciusyan 4/16/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue []*TreeNode

func (q *Queue) Push(root *TreeNode) {
	*q = append(*q, root)
}

func (q *Queue) Poll() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *Queue) Size() int {
	return len(*q)
}

func NewQueue() Queue {
	return Queue{}
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 10)
	// 层序遍历求解
	queue := NewQueue()
	queue.Push(root)

	for queue.Size() != 0 {
		size := queue.Size()
		// 一层一层遍历
		for i := 1; i <= size; i++ {
			node := queue.Poll()
			if i == size {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
	}

	return res
}
