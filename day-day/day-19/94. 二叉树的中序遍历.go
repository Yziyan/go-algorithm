// @Author: Ciusyan 10/5/23

package day_19

// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

// 递归实现版本
func inorderTraversal(root *TreeNode) []int {
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
