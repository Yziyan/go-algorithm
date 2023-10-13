// @Author: Ciusyan 10/13/23

package day_20

// IsCBT 判断是否是完全二叉树
func IsCBT(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 准备一个队列，用于层序遍历
	queue := NewQueue()
	queue.Offer(root)

	// 准备一个叶子节点的开关，如果开关被打开了，那么之后遇到的每一个节点，都必须是叶子节点
	leaf := false

	for queue.Size() != 0 {
		node := queue.Poll()

		// 处理当前节点
		if node.Left == nil && node.Right != nil {
			// 说明当前节点没有右对齐
			return false
		}

		if leaf && (node.Left != nil || node.Right != nil) {
			// 说明开关被打开了，但是没有满足条件
			return false
		}

		// 什么时候打开开关呢？
		if node.Left == nil || node.Right == nil {
			// 因为来到这里，已经不可能不是左对齐了
			leaf = true
		}

		// 处理左边
		if node.Left != nil {
			queue.Offer(node.Left)
		}

		// 处理右边
		if node.Right != nil {
			queue.Offer(node.Right)
		}
	}

	return true
}
