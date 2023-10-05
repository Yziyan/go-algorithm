// @Author: Ciusyan 10/5/23

package day_19

// https://leetcode.cn/problems/binary-tree-postorder-traversal/description/

// 递归实现版本
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)

	return res
}

// 后序遍历
func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 先对左右子树进行后序遍历
	postorder(root.Left, res)
	postorder(root.Right, res)
	// 搞事情
	*res = append(*res, root.Val)
}
