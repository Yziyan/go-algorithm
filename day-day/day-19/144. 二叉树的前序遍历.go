// @Author: Ciusyan 10/5/23

package day_19

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-preorder-traversal/

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	preorder(root, &res)

	return res
}

// 递归法实现的前序遍历
func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	// 先做事情
	*res = append(*res, root.Val)
	// 递归的前序遍历左右子树
	preorder(root.Left, res)
	preorder(root.Right, res)
}
