// @Author: Ciusyan 3/16/24

package phase_1

type Queue []*TreeNode

func NewQueue() Queue {
	return make(Queue, 0)
}

func (q *Queue) Push(val *TreeNode) {
	*q = append(*q, val)
}

func (q *Queue) Pop() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *Queue) Size() int {
	return len(*q)
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := NewQueue()
	queue.Push(root)

	for queue.Size() != 0 {
		size := queue.Size()

		curLevel := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Pop()
			curLevel[i] = node.Val

			if node.Left != nil {
				queue.Push(node.Left)
			}

			if node.Right != nil {
				queue.Push(node.Right)
			}
		}

		res = append(append([][]int(nil), curLevel), res...)
	}

	return res
}
