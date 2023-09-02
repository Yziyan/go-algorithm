// @Author: Ciusyan 2023/9/1

package day_2

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 准备一个队列，用于层序遍历
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	results := make([][]int, 0, 1)
	for len(queue) != 0 {
		// 取出当前层的节点数量
		size := len(queue)

		// 挨个将节点加入结果
		res := make([]int, 0, size)
		for i := 0; i < size; i++ {
			// 模拟出队
			node := queue[0]
			queue = queue[1:]

			res = append(res, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 添加到队列的前面
		results = append(append([][]int{}, res), results...)
	}

	return results
}
