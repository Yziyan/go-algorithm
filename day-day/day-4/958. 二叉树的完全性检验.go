// @Author: Ciusyan 2023/9/3

package day_4

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 准备一个队列，并且将根节点入队
	queue := make([]*TreeNode, 0, 1)
	queue = append(queue, root)

	// 叶子节点的开关
	leaf := false

	// 队列还有元素，就别停
	for len(queue) != 0 {
		// 先把当前层的节点个数取出来，一次操作完一层
		size := len(queue)
		for i := 0; i < size; i++ {
			// 将队头出队
			node := queue[0]
			queue = queue[1:]

			// 先看是否需要是 leaf，但是又不是 leaf
			if leaf && (node.Left != nil || node.Right != nil) {
				return false
			}

			// 操作左边
			if node.Left != nil {
				// 将其加入队列
				queue = append(queue, node.Left)
			} else if node.Right != nil {
				// 这里说明 左边是 nil，右边不是 nil，不满足左对齐
				return false
			}
			// 操作右边
			if node.Right != nil {
				// 将其加入队列
				queue = append(queue, node.Right)
			} else {
				// 来到这里说明不可能不是左对齐的情况了，可能是：
				//  1、左边是 nil，右边也是 nil
				//  2、左边不是 nil，右边也是 nil
				// 说明要打开 leaf 开关了
				leaf = true
			}
		}
	}

	// 遍历完都没返回 false，那就是完全二叉树
	return true
}
