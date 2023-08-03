// @Author: Ciusyan 2023/8/3

package tree

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 1)

	// 准备一个队列，用于层序遍历
	queue := make([]*TreeNode, 0)
	// 将根节点入队
	queue = append(queue, root)

	for len(queue) != 0 {
		// 取出当前层的节点数量
		size := len(queue)
		curLevel := make([]int, 0, size)
		for i := 0; i < size; i++ {
			// 模拟弹出队头元素
			node := queue[0]
			queue = queue[1:]

			// 记录结果
			curLevel = append(curLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		// 记录当前结果，放置开头
		res = append(append([][]int{}, curLevel), res...)
	}

	return res
}
