// @Author: Ciusyan 2023/9/1

package day_2

// https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	// 来到这里，说明两个节点都不为 nil，需要查看他们的值是否相等，还需要看：
	// 	左子树的左子树是否和右子树的右子树相等
	// 	左子树的右子树是否和右子树的左子树相等
	return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}
