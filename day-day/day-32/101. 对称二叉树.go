// @Author: Ciusyan 1/24/24

package day_32

// https://leetcode.cn/problems/symmetric-tree/description/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 递归含义是：left 和 right 是否是对称的
	var process func(left, right *TreeNode) bool
	process = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			// 如果左右都为 nil，肯定是对称的
			return true
		}

		if left != nil && right != nil {
			// 首先左右节点的值必须相等，其次是对应左右子树也必须是镜像的
			return left.Val == right.Val && process(left.Right, right.Left) && process(left.Left, right.Right)
		}

		// 能来到这里，说明左右只有一个为空树，肯定不是对称的
		return false
	}

	// 去看 root 的左、右 子树是否镜像
	return process(root.Left, root.Right)
}
