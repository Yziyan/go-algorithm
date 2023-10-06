// @Author: Ciusyan 10/6/23

package day_19

// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/

// 层序遍历，这个是一层一层的保存结果，是改进版的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 1)
	// 准备一个队列，并且将根节点入队
	queue := NewQueue()
	queue.Offer(root)

	// 只要队列不为 nil，就说明没有遍历完成
	for queue.Size() != 0 {
		// 要一层一层的收集，所以我们可以一层一层的弹出
		size := queue.Size()
		// 准备承接这一层的结果
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			// 弹出队头元素
			node := queue.Poll()
			temp[i] = node.Val

			// 看看有没有左右
			if node.Left != nil {
				queue.Offer(node.Left)
			}
			if node.Right != nil {
				queue.Offer(node.Right)
			}
		}

		// 加入这一层的结果
		res = append(res, temp)
	}

	return res
}

// 来个最经典的层序遍历，这个要求闭着眼睛都要会写
func level(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0, 1)
	// 准备一个队列，并且将根节点入队
	queue := NewQueue()
	queue.Offer(root)

	// 只要队列不为 nil，就说明没遍历完
	for queue.Size() != 0 {
		// 一个一个的弹，每次弹出一个
		node := queue.Poll()
		// 做事情
		res = append(res, node.Val)

		// 查看左右，有就入队
		if node.Left != nil {
			queue.Offer(node.Left)
		}
		if node.Right != nil {
			queue.Offer(node.Right)
		}
	}

	return res
}
