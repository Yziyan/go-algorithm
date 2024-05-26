// @Author: Ciusyan 5/26/24

package cycle_12_5_22_5_26

import "math"

// https://leetcode.cn/problems/balanced-binary-tree/

func isBalanced(root *TreeNode) bool {
	// 如果平衡，就返回这棵树的高度，不平衡就返回 -1
	var process func(root *TreeNode) float64
	process = func(root *TreeNode) float64 {
		if root == nil {
			return 0
		}

		l, r := process(root.Left), process(root.Right)

		if l == -1 || r == -1 {
			// 说明左子树或者右子树都已经不平衡了
			return -1
		}

		if math.Abs(l-r) > 1 {
			// 说明左右子树的高度差大于 1 了，违反了平衡树的性质
			return -1
		}

		// root 的高度就是，左子树和右子树的高度 + 1
		return max(l, r) + 1
	}

	return process(root) != -1
}
