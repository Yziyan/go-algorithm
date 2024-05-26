// @Author: Ciusyan 5/26/24

package cycle_13_5_22_5_26

import "math"

// https://leetcode.cn/problems/balanced-binary-tree/

/**
思路重复：
使用二叉树的递归套路即可求解，
如何判断一棵树是不是平衡二叉树呢？满足如下条件即可：
1.左子树是平衡二叉树
2.右子树是平衡二叉树
3.左右子树高度差 <= 1

所以我们的信息需要收集：
1.是否是平衡二叉树
2.树的高度

那么这俩属性可以合并到一起，如果树高是正常的，那么就返回正常的树高。
如果不是平衡的，那么就返回 -1 这样的高度即可。
 */


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
