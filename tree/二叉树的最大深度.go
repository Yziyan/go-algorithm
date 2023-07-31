// @Author: Ciusyan 2023/7/31

package tree

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左右子树的最大深度 + 1，就是整棵树的最大深度
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// 求解最大值
func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}
