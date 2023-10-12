// @Author: Ciusyan 10/12/23

package day_19

// https://leetcode.cn/problems/P5rCT8/

// 二叉树性质法
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	// 如果 p 节点有右子树，那么后继节点一定右子树上。
	succ := p.Right
	if succ != nil {
		// 说明有右子树，那么后继节点就存在于右子树的最左边
		for succ.Left != nil {
			succ = succ.Left
		}

		// 来到这里，肯定找到了后继
		return succ
	}

	// 来到这里，存在两种情况，1. 不存在后继 2. 存在于最祖先节点
	// 但是我们都可以统统从根节点往下遍历：

	cur := root
	for cur != nil {
		if p.Val > cur.Val {
			// 说明在右边还没找到比 p 大的，肯定还不可能是后继
			cur = cur.Right
		} else if p.Val == cur.Val {
			// 来到这，说明肯定找到了后继，可以不用遍历至末尾了
			break
		} else {
			// 说明有比 p 还大的了，可能是后继。
			succ = cur
			cur = cur.Left
		}
	}
	// 遍历完一条
	return succ
}

// 中序遍历法
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
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
