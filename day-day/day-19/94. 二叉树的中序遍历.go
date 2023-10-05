// @Author: Ciusyan 10/5/23

package day_19

// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

// 非递归实现
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	// 准备一个栈
	stack := NewStack()
	// 准备一个用于遍历的节点
	cur := root

	// 只要有一个不为 nil，就别停
	for cur != nil || stack.Size() != 0 {
		if cur != nil {
			// 我们需要一路向左，所以需要记录路径
			stack.Push(cur)
			// 然后向左走
			cur = cur.Left
		} else {
			// 来到这里，说明已经到了某一次的最左侧，需要弹出来访问
			node := stack.Pop()
			// 搞事情
			res = append(res, node.Val)
			// 然后将当前访问过的元素的右子节点当做根节点，看看它能否往下钻
			cur = node.Right
		}
	}

	return res
}

// 递归实现版本
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	inorder(root, &res)

	return res
}

// 中序遍历递归版
func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 先中序遍历左子树
	inorder(root.Left, res)
	// 搞事情
	*res = append(*res, root.Val)
	// 最后中序遍历右子树
	inorder(root.Right, res)
}
