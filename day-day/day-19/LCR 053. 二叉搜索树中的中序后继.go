// @Author: Ciusyan 10/12/23

package day_19

// https://leetcode.cn/problems/P5rCT8/

// 中序遍历法
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	var (
		succ *TreeNode
		pre  *TreeNode

		cur = root

		// 准备一个栈，用于中序遍历
		stack = NewStack()
	)

	// 只要有一个不为 nil，说明没遍历完
	for cur != nil || stack.Size() != 0 {

		if cur != nil {
			// 当前节点不为 nil，入栈
			stack.Push(cur)
			// 往左走
			cur = cur.Left
		} else {
			// 来到这里说明需要弹出栈顶
			node := stack.Pop()
			// 访问节点

			if pre == p {
				succ = node
				// 找到后继后，就可以返回了
				break
			}

			pre = node

			// 然后将 cur 变为栈顶元素的右子树
			cur = node.Right
		}
	}

	return succ
}
