// @Author: Ciusyan 10/13/23

package day_20

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/

// isCompleteTree 判断是否是完全二叉树，使用二叉树的递归套路版本
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type Info struct {
		isFull bool
		isCBT  bool
		height int
	}

	var getInfo func(root *TreeNode) *Info
	getInfo = func(root *TreeNode) *Info {
		if root == nil {
			return &Info{
				isFull: true,
				isCBT:  true,
				height: 0,
			}
		}

		// 收集左右子树的信息：
		leftInfo := getInfo(root.Left)
		rightInfo := getInfo(root.Right)

		var (
			// 构造信息，默认值怎么方便怎么来
			isFull = false
			isCBT  = false
			height = leftInfo.height
		)

		// 尝试右边推大 height
		if rightInfo.height > height {
			height = rightInfo.height
		}

		// 高的关系是否满足：左高 = 右高
		hr1 := leftInfo.height == rightInfo.height
		// 高的关系是否满足：左高 = 右高 + 1
		hr2 := leftInfo.height == rightInfo.height+1
		// 如何看满不满呢？左右都满，并且高度一致
		isFull = leftInfo.isFull && rightInfo.isFull && hr1

		// 讨论四种情况
		if isFull {
			// 满二叉树一定是完全二叉树
			isCBT = true
		} else if leftInfo.isFull && rightInfo.isFull && hr2 {
			// 左满、右满、左高 = 右高+1
			isCBT = true
		} else if leftInfo.isCBT && rightInfo.isFull && hr2 {
			// 左完全、右满、左高 = 右高+1
			isCBT = true
		} else if leftInfo.isFull && rightInfo.isCBT && hr1 {
			// 左满、右完全、左高 = 右高
			isCBT = true
		}

		return &Info{
			isFull: isFull,
			isCBT:  isCBT,
			height: height + 1,
		}
	}

	return getInfo(root).isCBT
}

// isCompleteTree1 判断是否是完全二叉树，使用层序遍历的版本
func isCompleteTree1(root *TreeNode) bool {
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
