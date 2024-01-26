// @Author: Ciusyan 1/26/24

package day_32

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 1)
	// 准备一个队列，用于层序遍历
	queue := NewQueue()
	queue.Add(root)

	direction := true // true 代表正向，false 代表反向
	for queue.Size() != 0 {
		// 此层要遍历多少个元素
		size := queue.Size()
		// 收集此层的答案
		temp := make([]int, size)
		if direction {
			// 说明是正向，那么正着收集结果
			for i := 0; i < size; i++ {
				// 弹出当前节点，并添加结果
				cur := queue.Remove()
				temp[i] = cur.Val

				// 然后依次添加 左和右
				if cur.Left != nil {
					queue.Add(cur.Left)
				}
				if cur.Right != nil {
					queue.Add(cur.Right)
				}
			}

		} else {
			// 说明是反向，那么反着收集结果
			for i := size - 1; i >= 0; i-- {
				// 弹出当前节点，并添加结果
				cur := queue.Remove()
				temp[i] = cur.Val

				// 然后依次添加 左和右
				if cur.Left != nil {
					queue.Add(cur.Left)
				}

				if cur.Right != nil {
					queue.Add(cur.Right)
				}
			}
		}

		// 保存此层的结果
		res = append(res, temp)
		// 转换方向
		direction = !direction
	}

	return res
}

type Queue []*TreeNode

func NewQueue() Queue {
	return make([]*TreeNode, 0)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Add(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Remove() *TreeNode {
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}
