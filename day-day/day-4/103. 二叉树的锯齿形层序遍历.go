// @Author: Ciusyan 2023/9/3

package day_4

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 1)
	// 准备一个队列，并将 root 添加进去
	queue := make([]*TreeNode, 1)
	queue[0] = root
	// 准备一个转向所使用的标记
	isReverse := false

	// 当队列不为空时，就别停
	for len(queue) != 0 {
		// 取出当前层的数量，一次性操作
		size := len(queue)

		curLevel := make([]int, 0, size)
		for i := 0; i < size; i++ {
			// 需要查看从哪边弹出
			if isReverse {
				// 需要反着来，就从队尾弹出
				last := len(queue) - 1
				node := queue[last]
				queue = queue[:last]

				curLevel = append(curLevel, node.Val)

				// 然后查看右左子树，还有没有子节点，需要注意：
				//	1、先看右子树 2、添加到队列的头部
				if node.Right != nil {
					queue = append([]*TreeNode{node.Right}, queue...)
				}

				if node.Left != nil {
					queue = append([]*TreeNode{node.Left}, queue...)
				}
			} else {
				// 正过来了
				node := queue[0]
				queue = queue[1:]

				curLevel = append(curLevel, node.Val)

				// 正常层序遍历的判断
				if node.Left != nil {
					queue = append(queue, node.Left)
				}

				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}

		}

		// 加入结果后，别忘了改变方向
		res = append(res, curLevel)
		isReverse = !isReverse
	}

	return res
}
