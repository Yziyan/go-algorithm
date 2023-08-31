// @Author: Ciusyan 2023/9/1

package day_2

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 求出做右子树的最大深度，然后 +1，就是这棵树的最大深度
	lD := maxDepth(root.Left)
	rD := maxDepth(root.Right)

	if lD > rD {
		return lD + 1
	}

	return rD + 1
}
