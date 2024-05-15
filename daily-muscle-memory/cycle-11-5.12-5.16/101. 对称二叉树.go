// @Author: Ciusyan 5/16/24

package cycle_11_5_12_5_16

// https://leetcode.cn/problems/symmetric-tree/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var process func(left, right *TreeNode) bool
	process = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && process(left.Right, right.Left) && process(left.Left, right.Right)
	}

	return process(root, root)
}
