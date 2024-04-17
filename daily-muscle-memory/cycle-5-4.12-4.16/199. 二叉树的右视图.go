// @Author: Ciusyan 4/16/24

package cycle_5_4_12_4_16

// https://leetcode.cn/problems/binary-tree-right-side-view/

/**
思路重复：
对于这个题，要收集左视图或者右视图，其实就是访问到每一层的第一个节点，或者最后一个节点。
对于这题，就是想访问每一层的最后一个节点，
那么我们进行层序遍历，并且一层一层的进行遍历。
当遍历到每一层的末尾节点的时候，就将答案收集起来。
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := make([]*TreeNode, 0, 100)
	queue = append(queue, root)
	res := make([]int, 0, 100)

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				// 说明到达了最后一个节点
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 前面的 size 个，已经遍历完成了，直接弹出
		queue = queue[size:]
	}

	return res
}

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

func rightSideView2(root *TreeNode) []int {
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
